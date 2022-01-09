# Liv팀에서 만든 WebBrowser를 이용한 화면속 색상찾기
## 본 프로그램은 1번 모니터의 색상만 감지합니다.

빌드전 필요한목록 [설치해야함.]
1. Golang 1.5 버전이상
2. TDM-GCC-64 가급적 최신버전 (https://jmeubank.github.io/tdm-gcc/)
3. Cmake 3.16 버전이상.

빌드 방법
1. 프로젝트 소스코드를 다운로드 하고 압축해제
2. 압축 해제된 폴더에서 powershell 혹은 cmd를 실행하여 (go mod tidy 입력) 
3. 압축해제된 폴더에 있는 Build.bat 실행 (exe 파일이 생성 안되면 Powershell로 실행후 오류 내용확인하기)
### ※ build.bat을 실행했지만 exe파일이 생성이 안되는경우
### win + r 눌러서 실행창에 powershell 치고 엔터 압축해제된 폴더의 주소를 복사후 powershell창에 cd 입력후 ctrl + v 후 엔터
### go build ./ 입력
### ※ 오류가 있을경우 github의 Issue에 적어주시면 답변 드리겠습니다.
