<template>
  <div class="row items-push">
    <div class="col-sm-6 col-md-4 col-xl-3 d-flex align-items-center" v-for="skin in skins" :key="skin.name">
      <div class="options-container w-100">
        <div class="options-item block block-rounded block-transparent mb-0">
          <div class="block-content text-center">
            <p class="mb-2 overflow-hidden">
              <img class="img-fluid" :src="assetsURL(skin.cover_imgs['index.jpg'])" alt />
            </p>
            <p class="font-w600 mb-0">{{ skin.name }}</p>
            <p class="font-size-sm text-muted">{{ skin.description }}</p>
          </div>
        </div>
        <div class="options-overlay rounded-lg bg-white-50">
          <div class="options-overlay-content">
            <div class="mb-3">
              <a class="btn btn-hero-light" href="javascript:void(0)">
                <i class="fa fa-eye text-primary mr-1"></i> View
              </a>
            </div>
            <div class="btn-group">
              <a class="btn btn-sm btn-light" href="javascript:void(0)">
                <i class="fa fa-download text-black mr-1"></i>
              </a>
              <a class="btn btn-sm btn-light" href="javascript:void(0)">
                <i class="fa fa-trash text-danger mr-1"></i>
              </a>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
<script>
import { getAllSkin } from "@/api/skin";

export default {
  data () {
    return {
      skins: []
    }
  },
  mounted() {
    getAllSkin().then(res => {
      console.log(res);
      this.skins = res.skin
    });
  }
};
</script>