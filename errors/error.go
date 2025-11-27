// 파일 경로: apperror/errors.go (패키지 이름 변경)
package planet_err

import "errors"

// 닉네임 중복을 나타내는 사용자 정의 오류 타입
var ErrNicknameDuplicate = errors.New("nickname is already in use")
