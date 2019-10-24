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