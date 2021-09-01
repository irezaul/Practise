## TailwindCss Tutorial

### Download [node.js](https://nodejs.org/en/) on your pc -

1. open your working directoy 

> use this command
```bash
npm init -y
```
2nd Commnad use it for install tailwindcss -
```bash
npm install tailwindcss
```

#### Now create folder name `src` & under create `styles.css` file & paste 
![src](https://user-images.githubusercontent.com/77927449/131243735-977eb7f3-6871-4f89-a497-8a84df8affc0.jpg)


#### Include Tailwind in your CSS
> Create a CSS file if you don’t already have one, and use the `@tailwind directive` to inject `Tailwind’s base, components,` and `utilities` styles:

> now paste the `tailwindcss` code on style.css src folder -

```bash
@tailwind base;
@tailwind components;
@tailwind utilities;
```
## Now open package.json file & edit the line & write this
```bash
"scripts": {
    "build-css": "tailwindcss build src/styles.css -o public/styles.css"
```
![package_json](https://user-images.githubusercontent.com/77927449/131676447-6b8f55e2-e910-4972-a79b-8eee5cc882f5.jpg)
























