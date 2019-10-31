# Firebase jwt authorization

---

クライアント(testApp)は、jwt認可と重要なデータを扱うjwt認証機能を持つセーバー(src/main.go)にgRPCでアクセスし、ユーザー情報などを取得する。
jwt認可は、jwtを入力とし、Firebaseによりjwtが正しいことを確認し、base64デコードによりユーザー情報を取得する。
一方、ここで言うjwt認証は、認証機能があるわけではなく、jwtによる認可を2度行い、二回目のjwtは再度サインインを行ったタイミングで新たに生成されるワンタイムjwtを用いることで、jwtが盗まれている場合でも情報を保護できるようにしている。

Firebaseにアクセスして認可を行うサーバー機能は、src/jwtAuthServerに集められたjwtauthserverパッケージである。
また、gRPCのクライアントとなる機能は、src/clientのjwtauthclientパッケージで、testAppはこのパッケージを用いて構築されている。

クライアントはURLルーティング機能を持ち、https://github.com/k-washi/example-vue-cli からのクライアントからアクセスできる。

```yaml
- path: GET: "/jwt/ex-jwt-auth"
  info: "jwt認可"
- path: GET: "/auth/ex-authentication"
  info: "jwt認証"

```

"/auth/ex-authentication"は、サーバ側にすでに一度目のjwtが保存されている場合、2度目のjwtとして扱われ200を返す。一方で、サーバーにjwtが保存されていない場合、一度目のjwtとしてサーバーに保存され、サインイン後のjwtを送るよう202を返す。



## Setup

Firebaseのサービスアカウント（Firebase Admin SDK)より秘密鍵ファイルを生成する。
また、ファイルのパスもk8sで環境変数として設定する。

```bash
export GOOGLE_APPLICATION_CREDENTIALS="/home/user/Downloads/service-account-file.json"
```

## proto Setup

```bash
src/jwtAuthpb/generate.sh
```

## start

start server

```bash
go run src/main.go
```

start test app client 

```bash
go run testApp/main.go
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