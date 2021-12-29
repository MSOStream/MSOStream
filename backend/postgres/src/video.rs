use crate::*;
use models::{NewVideo, Video};

pub fn create_video<'a>(
    conn: &PgConnection,
    id: &'a str,
    user_id: &'a str,
    title: &'a str,
    description: &'a str,
) -> usize {
    use schema::videos;

    let new_video = NewVideo {
        id,
        user_id,
        title,
        description,
    };

    let result = diesel::insert_into(videos::table)
        .values(&new_video)
        .execute(conn)
        .expect("Error saving new video");
    result
}

pub fn list_videos(conn: &mut PgConnection) -> Vec<Video> {
    schema::videos::table
        .load::<Video>(conn)
        .expect("Could not load videos")
}
