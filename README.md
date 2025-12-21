# planet_utils

gRPC 컴파일 명령 : protoc --proto_path=./protos --go_out=. --go-grpc_out=. ./protos/*proto

`protoc` 명령어 옵션들의 역할과 설명

| 명령어 / 옵션 | 역할 | 설명 |
|---------------|------|------|
| `protoc` | 컴파일러 실행 | Protobuf 컴파일러 바이너리를 실행합니다. |
| `--proto_path=./protos` | 입력 파일 경로 | `.proto` 파일이 위치한 디렉토리를 지정합니다. `import` 경로 해석에도 사용됩니다. |
| `--go_out=./pb` | Go 코드 출력 설정 | 일반 Protobuf 메시지(구조체) 코드를 `./pb` 폴더에 생성합니다. |
| `--go-grpc_out=./pb` | gRPC 코드 출력 설정 | gRPC 서버/클라이언트 인터페이스 코드를 `./pb` 폴더에 생성합니다. |
| `./protos/*.proto` | 컴파일 대상 파일 | `./protos/` 디렉토리 내부의 모든 `.proto` 파일을 컴파일 대상으로 지정합니다. |


## 브랜치 생성
git checkout -b [브랜치 명]
## 브랜치 등록
git push -u origin [브랜치 명]