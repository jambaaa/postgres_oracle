```
#   **Connect Oracle Database from GoLang service**
# **1.Go DRiver for ORacle installation** 
## Ажиллаж буй folder-т доорх командыг хуулж import хийнэ.
    go get github.com/godror/godror@latest 
## Connection
    db, err := sql.Open("godror", `user="your_user_name" password="your_password" connectString="host:port/service_name"`)

Дэлгэрэнгүй : https://github.com/godror/godror
# **2.Install Oracle Instant Client on Ubuntu 20.04**   
## **Step 1. Update Ubuntu**
    sudo apt update && sudo apt upgrade
## **Step 2. Download Required Packages**
##  [Instant-client-package](https://download.oracle.com/otn_software/linux/instantclient/19800/instantclient-basic-linux.x64-19.8.0.0.0dbru.zip?xd_co_f=fae07a5a-ae5f-4c1a-8ce0-f41ed4a394c5 "Instant Client package")

##  [Sqlplus-package](https://download.oracle.com/otn_software/linux/instantclient/19800/instantclient-sqlplus-linux.x64-19.8.0.0.0dbru.zip "sqlplus package" )

## **Step 3. Extract the Packages**
##  1.Opt дотор oracle folder шинээр үүсгэнэ.
        sudo mkdir /opt/oracle
##  2.Download хийж авсан 2 zip file-ийг opt/oracle/ folder-т unzip хийх
        /Downloads

        sudo unzip instantclient-basic-linux.x64-19.8.0.0.0dbru.zip -d /opt/oracle
        sudo unzip instantclient-sqlplus-linux.x64-19.8.0.0.0dbru.zip -d /opt/oracle/
## **Step 4. Configure SQLPlus**
##  1.Terminal дээр доорх командыг бич.
        nano ~/.profile
##  2.Дараах 2 мөрийг хамгийн сүүлд нь хуул. 
        export PATH="$PATH:/opt/oracle/instantclient_19_8"
        export LD_LIBRARY_PATH="$LD_LIBRARY_PATH:/opt/oracle/instantclient_19_8"
        sudo sh -c "echo /opt/oracle/instantclient_19_8 > /etc/ld.so.conf.d/oracle-instantclient.conf"
        sudo ldconfig
##  3.Дээрх файлыг хадгалаад хаах.Доор командыг дахин ажилуул.
        source ~/.profile

Дэлгэрэнгүй :  https://manjaro.site/how-to-install-oracle-instant-client-on-ubuntu-20-04/#google_vignette
# **3. Install the operating system libaio package with sudo or as the root user. For example:**   
        sudo apt-get install libaio1
Дэлгэрэнгүй : https://oracle.github.io/odpi/doc/installation.html#linux

```