use redis::{
    Client, Commands, Connection, ErrorKind, FromRedisValue, RedisError, RedisResult, ToRedisArgs,
};
use rocket::http::uri::Origin;
use rocket::outcome::Outcome;
use rocket::request::{self, FromRequest};
use rocket::Request;
use std::env;

pub struct RedisCache {
    client: Client,
}

impl RedisCache {
    fn get_connection(&self) -> Connection {
        self.client
            .get_connection()
            .expect("Unable to establish redis connection")
    }

    pub fn new() -> Self {
        RedisCache {
            client: Client::open(env::var("REDIS_URL").expect("REDIS_URL must be set"))
                .expect("Can't establish redis connection"),
        }
    }
}

fn uri_to_string(origin: &Origin) -> String {
    let base = origin.path().as_str().to_owned();
    let params = origin.query().map_or_else(|| "", |q| q.as_str());
    if params.is_empty() {
        return base;
    }

    base + "?" + params
}

pub struct CacheHelper {
    cache: RedisCache,
    key: Option<String>,
}

impl CacheHelper {
    fn new(key: Option<String>) -> Self {
        CacheHelper {
            cache: RedisCache::new(),
            key,
        }
    }

    #[allow(dead_code)]
    fn set_key(&mut self, key: String) {
        self.key = Some(key);
    }

    pub fn set_cache<T: ToRedisArgs>(self, value: T) {
        if self.key.is_none() {
            // TODO: Actual error handling
            return;
        }

        let key = self.key.unwrap();

        let _: () = self
            .cache
            .get_connection()
            .set(key, value)
            .expect("Unable to cache");
    }

    // TODO: beautify this
    pub fn get_cache<T: FromRedisValue>(&self) -> RedisResult<T> {
        let key = self
            .key
            .as_ref()
            .ok_or_else(|| RedisError::from((ErrorKind::InvalidClientConfig, "key isn't set")))?;
        self.cache.get_connection().get(key)
    }

    pub fn del_cache<K: ToRedisArgs>(&self, key: K) {
        let _: RedisResult<()> = self.cache.get_connection().del(key);
    }
}

impl Default for CacheHelper {
    fn default() -> Self {
        CacheHelper::new(None)
    }
}

#[derive(Debug)]
pub struct CacheError;

#[rocket::async_trait]
impl<'r> FromRequest<'r> for CacheHelper {
    type Error = CacheError;

    async fn from_request(req: &'r Request<'_>) -> request::Outcome<Self, Self::Error> {
        let cache_helper = CacheHelper::new(Some(uri_to_string(req.uri())));
        Outcome::Success(cache_helper)
    }
}
