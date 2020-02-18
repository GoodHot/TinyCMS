<template>
  <div class="t-publish-advanced">
    <div class="row justify-content-center">
      <div class="col-md-8 col-lg-6">
        <form
          action="be_pages_ecom_product_edit.html"
          method="POST"
          onsubmit="return false;"
          data-select2-id="9"
        >
          <div class="form-group">
            <label for="one-ecom-product-meta-description">文章简介</label>
            <textarea
              class="js-maxlength form-control js-maxlength-enabled"
              id="one-ecom-product-meta-description"
              name="one-ecom-product-meta-description"
              rows="4"
              maxlength="115"
              data-always-show="true"
              data-placement="top"
            >Dark Souls III is an action role-playing video game developed by FromSoftware.</textarea>
            <small class="form-text text-muted">115 Character Max</small>
          </div>
        </form>
        <form class="dropzone dz-clickable" action="be_pages_ecom_product_edit.html">
          <div class="dz-default dz-message">
            <span>设置封面图</span>
          </div>
        </form>
      </div>
    </div>
  </div>
</template>