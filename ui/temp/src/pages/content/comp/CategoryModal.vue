<template>
  <b-form>
    <b-form-group label="分类名称" label-for="input-1" >
      <b-form-input placeholder="分类名称"></b-form-input>
    </b-form-group>

    <b-form-group label="分类标题" label-for="input-1">
      <b-form-input id="input-2" required placeholder="分类标题"></b-form-input>
    </b-form-group>

    <b-form-group label="SEO标题" label-for="input-1">
      <b-form-input id="input-2" required placeholder="SEO标题"></b-form-input>
    </b-form-group>

    <b-form-group label="SEO关键字" label-for="input-1">
      <b-form-input id="input-2" required placeholder="SEO关键字"></b-form-input>
    </b-form-group>

    <b-form-group label="SEO描述" label-for="input-1">
      <b-form-textarea placeholder="SEO描述" rows="3" max-rows="6"></b-form-textarea>
    </b-form-group>

    <b-form-group label="分类图标" label-for="input-1">
      <b-form-file
        placeholder="上传图片"
        drop-placeholder="Drop file here..."
      ></b-form-file>
    </b-form-group>
  </b-form>
</template>