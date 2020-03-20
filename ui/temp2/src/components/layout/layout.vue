<template>
  <div class="t-layout-container sidebar-o enable-page-overlay side-scroll page-header-fixed page-header-dark main-content-narrow">
    <aside class="t-layout-side">
      side-overlay
      <slot name="side"></slot>
    </aside>
    <nav>
      Navigation
      <slot name="nav"></slot>
    </nav>
     <header>
       header
       <slot name="header"></slot>
     </header>
     <main>
       main
       <slot></slot>
     </main>
     <footer>
       <slot name="footer"></slot>
     </footer>
  </div>
</template>