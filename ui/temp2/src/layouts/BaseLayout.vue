<template>
  <div id="page-container" class="sidebar-o enable-page-overlay side-scroll page-header-fixed page-header-dark main-content-narrow">
    <Nav></Nav>
    <Header></Header>
    <main id="main-container">
      <div class="content content-full bg-body-dark" >
        <router-view />
      </div>
    </main>
    <Footer></Footer>
  </div>
</template>
<script>
import Nav from './BaseLayoutComp/Nav'
import Header from './BaseLayoutComp/Header'
import Footer from './BaseLayoutComp/Footer'
export default {
  components: {
    Nav,
    Header,
    Footer
  }
}
</script>