# schoolMenu-API
/[학교명] 으로 심플하고 빠른 급식 API 학교 급식 API

## 빠른 사용법
오늘의 메뉴를 불러옵니다.
> ###### 학교명은 3글자 이상 입력하는 것을 추천합니다
```
https://schul.ml/경북소프트웨어고등학교   or   https://schul.ml/경북소
```

## 추가 사용법

### 특정 날짜의 급식 가져오기
```
https://schul.ml/경북소/20220103
```
### 같은 이름의 학교가 여러개일 때
아래 주소로 학교를 검색하여 원하는 학교의 교육청코드(apt_code)를 찾습니다. (JSON beautifier 사용추천..)
```
https://schul.ml/school/[학교명]
```
그리고 다음과 같은 형식으로 오늘의 급식을 불러올 수 있습니다.
```
https://schul.ml/경북소/today/[교육청코드 (예시: R10)]Cancel changes
```
응답 형식
---

### 급식 검색 응답
요청:
```
https://schul.ml/경북소
```
응답 (검색된 학교가 하나일때):
```
{
  "status": {
    "success": true,
    "msg": "success",
    "server_date": "2022-01-03",
    "searched_school": "경북소프트웨어고등학교",
    "school_apt_name": "경상북도교육청"
  },
  "menu_date": "20220103",
  "breakfast": [
    "쌀밥(조)",
    ...
  ],
  "lunch": [
    "쌀밥(조)",
    ...
  ],
  "dinner": [
    "발아현미밥",
    ...
  ]
}
```
응답 (검색된 학교가 두 개 이상일때):
```
{
  "status": {
    "success": true,
    "msg": "success | 학교가 두 개 이상 검색되었습니다. /[학교명]/today/[교육청코드] 로 다른학교도 검색해보세요. 교육청코드(apt_code)는 /school/[학교명] 에서 확인할 수 있습니다",
    "server_date": "2022-01-03",
    "searched_school": "장곡중학교",
    "school_apt_name": "경기도교육청"
  },
  "menu_date": "20220103",
  "breakfast": null,
  "lunch": [
    "찹쌀밥",
    ...
  ],
  "dinner": null
}
```
### 학교 검색 응답
요청:
```
https://schul.ml/school/경북소
```
응답:
```
{
  "status": {
    "success": true,
    "msg": "success",
    "server_date": "2022-01-03",
    "searched_school": "경북소프트웨어고등학교",
    "school_apt_name": ""
  },
  "list": [
    {
      "apt_code": "R10",
      "apt_name": "경상북도교육청",
      "school_code": "8750767",
      "school_name": "경북소프트웨어고등학교"
    },
  ]
}
```

## 개요
### " 누구나 편리하게 쓸 수 있도록 만들자 "
neis API에서는 학교의 급식을 검색할 때 학교가 소속된 교육청의 코드와 학교 고유 코드가 있어야 검색이 가능합니다.  
하지만 매번 학교의 교육청 코드와 고유 코드를 찾아보기는 어렵습니다.  
따라서 만든것이 이 프로젝트 입니다.  
  
사용자 편의를 우선으로 학교명으로만 간단하게 학교 급식을 가져올 수 있도록 만들었습니다

## 향후 계획
현재로썬 Go 를 배운지 많이 안됐으므로 계속 배워가며 코드의 퀄리티를 높이고, 현재보다 더 빠르고 가볍게 코드를 업그레이드 할 계획입니다.
