# GOLANG JWT RSA STRATEGY

## GENERATE PRIVATE KEY

```cmd
ssh-keygen -t rsa -f keys/key.pem -m pkcs8 -C "demo@example.com"
```

- output:

```cmd
Generating public/private rsa key pair.
Enter passphrase (empty for no passphrase):
Enter same passphrase again:
Your identification has been saved in keys/key.pem.
Your public key has been saved in keys/key.pem.pub.
The key fingerprint is:
SHA256:O4jysIohNoO+tSF3PcoOpgwUTqZDMrW62vFakoQoe40 demo@example.com
The key's randomart image is:
+---[RSA 3072]----+
|  .              |
| . .             |
|o=.              |
|O+.              |
|Bo.     S        |
|+= + ... .       |
|B=E*+..oo        |
|O+OX* . ..       |
|+*=o++           |
+----[SHA256]-----+
```

## GENERATE PUBLIC KEY

```cmd
ssh-keygen -t rsa -f keys/key.pem -e -m pkcs8 -C "demo@example.com" > keys/key.pem.pub
```

