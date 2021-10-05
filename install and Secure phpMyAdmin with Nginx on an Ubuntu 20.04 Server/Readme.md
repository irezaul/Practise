 ### how to install and secure phpMyAdmin with Nginx on an Ubuntu 20.04 server.
 
 ### 1. update apt 
 ```bash
 sudo apt update
 ```
### 2. install nginx
 
 ```bash
 sudo apt install nginx
 ```
 ![install nginx](https://user-images.githubusercontent.com/77927449/135949819-a2f93423-1b19-4fd3-8130-c74340fe7f6d.png)

### 3. install mysql-server

```bash
apt install mysql-server
```
![install mysql-server](https://user-images.githubusercontent.com/77927449/135950092-d212533a-e421-46b3-bcd7-6baca8f128a1.png)

### 4. install mysql_sequre_installation

```bash
mysql_sequre_installation
```
![mysql_sequre_installation](https://user-images.githubusercontent.com/77927449/135950108-77cbf4f9-c0a3-4b57-9ce5-523ce8ee1210.png)

![mysql secure done](https://user-images.githubusercontent.com/77927449/135950317-49a8d0f7-368c-4c7d-aca3-13f8e5ec72f6.png)

### 5. install php-fpm & php-mysql
```bash
apt install php-fpm php-mysql
```
![php-fpm php-mysql](https://user-images.githubusercontent.com/77927449/135950500-a44d2641-bbda-4952-9d58-aa87c4e15c88.png)

#### 6.check & mysql version

![php   mysql version](https://user-images.githubusercontent.com/77927449/135951541-a7671020-8e13-4f31-9764-dc3670dab4a5.png)

#### 7. now go to Nginx folder & check the html root directory also you can edit your html file.

![nginx folder](https://user-images.githubusercontent.com/77927449/135951026-fdaf6dbb-ac7f-41b5-8d17-a7b84312a35f.png)

### 8. root check
![html root](https://user-images.githubusercontent.com/77927449/135951067-f969ce73-050a-4987-bb38-9ecdcf7f084b.png)

![root 1](https://user-images.githubusercontent.com/77927449/135951097-27339cd8-de2c-4bf7-a0b8-16834056365a.png)

![root 2](https://user-images.githubusercontent.com/77927449/135951114-41cb3e48-c521-4fbd-9422-accbe58a516d.png)

#### 9. install sucess Nginx
![nginx install sucess](https://user-images.githubusercontent.com/77927449/135951160-a2ba15e6-697b-49c8-9863-347d4713475e.png)

#### •• edit the hmtl
![master academy](https://user-images.githubusercontent.com/77927449/135951143-07963b63-2660-4879-9505-c91caba5ad98.png)




