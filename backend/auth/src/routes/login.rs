use crate::construct_response;
use crate::storage::Store;
use crate::user::User;
use crate::users::v1::{LoginRequest, LoginResponse};
use opentelemetry::trace::TraceContextExt;
use tonic::{Response, Status};
use tracing_opentelemetry::OpenTelemetrySpanExt;

#[tracing::instrument]
pub async fn handle_login_request(req: LoginRequest) -> Result<Response<LoginResponse>, Status> {
    let id = tracing::Span::current()
        .context()
        .span()
        .span_context()
        .trace_id();
    tracing::info!("handling login request, TraceId: {}", id);
    let mut store = Store::new();

    let user = store.get_user(&req.email).into();

    verify_password(&user, req.password.as_bytes())
}

#[tracing::instrument]
fn verify_password(user: &User, pwd: &[u8]) -> Result<Response<LoginResponse>, Status> {
    match argon2::verify_encoded(&user.password, pwd) {
        Ok(true) => {
            let mut resp = Response::new(LoginResponse {});
            construct_response(&mut resp, &user.id);
            Ok(resp)
        }
        Ok(false) => Err(Status::invalid_argument("Wrong password")),
        Err(_) => Err(Status::internal("Internal hashing error")),
    }
}
