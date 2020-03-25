<template>
  <div id="page-container">
    <main id="main-container">
      <div class="bg-image" :style="`background-image: url('${bgImg}');`">
        <div class="row no-gutters justify-content-center bg-primary-dark-op">
          <div class="hero-static col-sm-8 col-md-6 col-xl-4 d-flex align-items-center p-2 px-sm-0">
            <router-view />
          </div>
        </div>
      </div>
    </main>
  </div>
</template>
<script>
import bgImg from '@/assets/imgs/auth_bg.jpg'

export default {
  data () {
    return {
      bgImg
    }
  }
}
</script>