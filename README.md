# setting GOPATH for mac

- zshell -> touch ~/.zshrc -> open -e ~/.zshrc -> CONFIG -> source ~/.zshrc
- bash -> touch ~/.bash_profile -> open -e ~/.bash_profile -> CONFIG -> source ~/.bash_profile
- CONFIG:\
   export GOPATH=/Users/$USER/go
export PATH=$GOPATH/bin:$PATH

# install gin livereload

- which go
- cd $GOPATH/bin/
- go install github.com/codegangsta/gin@latest
- go mod tidy
- gin -p 8500 -a 8005 run main.go

# set .env

```
DB_HOST=
DB_PORT=
DB_USER=
DB_PWD=x
DB_NAME=
DBL_HOST=
DBL_PORT=
DBL_USER=
DBL_PWD=
DBL_NAME=
DBR_HOST=
DBR_PORT=
DBR_USER=
DBR_PWD=
DBR_NAME=
DBLR_HOST=
DBLR_PORT=
DBLR_USER=
DBLR_PWD=
DBLR_NAME=
JWT_RF_KEY=
JWT_AC_KEY=
AWS_BUCKET=
AWS_REGION=ap-southeast-1
AWS_ACCESS_KEY_ID=
AWS_SECRET_ACCESS_KEY=
SMS_API_KEY=
SMS_API_SECRET_KEY=
SMS_SENDER=APSTH
TK_PUBPLIC_KEY=
EMAIL_NAME=admin@app-apsx.com
EMAIL_PWD=
BASE_URL=https://www.clinic.app-apsx.com
JWT_RF_EXPIRE=720
JWT_AC_EXPIRE=6
API_PORT=:8005
TZ=Asia/Bangkok
ENV=DEV


```
