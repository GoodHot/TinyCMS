<template>
  <div class="row no-gutters flex-md-10-auto">
    <div class="col-md-4 col-lg-5 col-xl-3">
      <div class="content">
        <div class="d-md-none push">
          <button
            type="button"
            class="btn btn-block btn-hero-primary"
          >Side Content</button>
        </div>
        <div id="side-content" class="d-none d-md-block push">
          <p>Your side content..</p>
        </div>
      </div>
    </div>
    <div class="col-md-8 col-lg-7 col-xl-9 bg-body-dark">
      <div class="content content-full">
        <div class="block block-fx-pop">
          <div class="block-content">
            <p>Your main content..</p>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
<script>
export default {
  data() {
    return {
      contentSide: true
    };
  }
};
</script>