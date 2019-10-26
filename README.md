# Firebase jwt authorization

---

## install

```bash
go get firebase.google.com/go
go get -u google.golang.org/api/option
```

Firebaseのサービスアカウント（Firebase Admin SDK)より秘密鍵ファイルを生成する。
このファイルは公開せず、k8sのConfigMapで渡すものとする。

また、ファイルのパスもk8sで環境変数として設定する。

```bash
export GOOGLE_APPLICATION_CREDENTIALS="/home/user/Downloads/service-account-file.json"
```

## proto Setup

```bash
src/jwtAuthpb/generate.sh
```

## Firebase jwt payload

```json
//Payload
{
  "name": "testuser",
  "iss": "https://securetoken.google.com/ex-firebase-auth",
  "aud": "ex-firebase-auth",
  "auth_time": 1572007184,
  "user_id": "qZhsF2HfuWZEBghFa4nl2Kidyp22",
  "sub": "qZhsF2HfuWZEBghFa4nl2Kidyp22",
  "iat": 1572007184,
  "exp": 1572010784,
  "email": "test@test.com",
  "email_verified": false,
  "firebase": {
    "identities": {
      "email": [
        "test@test.com"
      ]
    },
    "sign_in_provider": "password"
  }
}

```