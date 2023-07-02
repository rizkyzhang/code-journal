<!-- Reference:
https://dilshankelsen.com/customize-input-file-button/
 -->

```html
<input type="file" id="upload-file" hidden />
<label for="upload-file">Upload</label>
<!-- OR -->
<label>
  <input type="file" hidden />
  Upload
</label>
```

```css
input[type="file"]::file-selector-button {
  /* Add properties here */
}
```
