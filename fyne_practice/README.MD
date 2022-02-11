

### if error when build or go run use this method because Some Linux systems lack the compile headers for graphics libraries that are installed. -

![](./img/build_error.png)

> now use command for gcc libg install
```bash
sudo apt-get install golang gcc libgl1-mesa-dev xorg-dev
```

### #01 Create a input with placeholder

```go
  Variable := widget.NewEntry()
	Variable.PlaceHolder = "PlaceHolder Text.."
  ```
