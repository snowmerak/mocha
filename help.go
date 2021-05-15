package main

var helpPhrase = `Mocha Bot 설명서
- script
1. Go 스크립트 실행
Me:
run go
"hello, " + "mocha"
----------
mocha:
hello, mocha

1-1. Go 스크립트 실행 시간 측정
Me:
estimate go
"hello, mocha"
----------
mocha:
hello, mocha
297.123µs

2. JS 실행
Me:
run js
5 + 8
----------
mocha:
13

2-2. JS 실행 시간 측정
Me:
run js
5 + 8
----------
mocha:
13
430.385µs

- problem
1. 랜덤 문제
Me:
random problem
----------
mocha:
제목: <<problem title>>
설명: <<problem explanation>>

2. 문제 보기
Me:
view problem
<<problem title>>
----------
mocha:
제목: <<problem title>>
설명: <<problem explanation>>

3. 답안 제출
Me:
submit solution
<<problem title>>
<<solution code>>
----------
mocha:
<<succeed or faild>>

답안 제출 시 solution 함수를 작성해서 제출해주세요.

4.문제 만들기
Me:
enroll problem
name
<<problem title>>
explain
<<problem explanation>>
cases
<<nth test case input...>>
<<nth test case output...>>
----------
mocha:
<<succeed or faild>>

테스트 케이스의 경우 입출력 모두 공백 하나를 기준으로 작성하고 검사할 땐 모든 원소가 ', '로 연결되어 검사됩니다.
출력이 1 2 3 으로 지정되어 있다면 출력 검사는 1, 2, 3 이 맞는 지 검사합니다.
즉, solution 함수의 출력이 문자열, "1, 2, 3" 이어야 합니다.
`
