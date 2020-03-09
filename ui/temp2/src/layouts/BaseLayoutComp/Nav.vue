<template>
  <nav id="sidebar" aria-label="Main Navigation">
    <div class="bg-header-dark">
      <div class="content-header bg-white-10">
        <a class="link-fx font-w700 font-size-lg text-white" href="index.html">
          <span class="smini-visible">
            <span class="text-white-75">
              <i class="fab fa-wordpress"></i>
            </span>
          </span>
          <span class="smini-hidden">
            <span class="text-white">TinyCMS</span>
          </span>
        </a>
        <div>
          <a
            class="d-lg-none text-white ml-2"
            data-toggle="layout"
            data-action="sidebar_close"
            href="javascript:void(0)"
          >
            <i class="fa fa-times-circle"></i>
          </a>
        </div>
      </div>
    </div>
    <div class="content-side content-side-full">
      <ul class="nav-main">
        <li class="nav-main-item open">
          <a class="nav-main-link nav-main-link-submenu" href="#">
            <t-icon class="nav-main-link-icon" icon="campground" />
            <span class="nav-main-link-name">控件</span>
          </a>
          <ul class="nav-main-submenu">
            <li class="nav-main-item">
              <a class="nav-main-link" href>
                <span class="nav-main-link-name">图标</span>
              </a>
            </li>
          </ul>
        </li>
        <li class="nav-main-item open">
          <a
            class="nav-main-link nav-main-link-submenu"
            data-toggle="submenu"
            aria-haspopup="true"
            aria-expanded="true"
            href="#"
          >
            <t-icon class="nav-main-link-icon" icon="car-battery" />
            <span class="nav-main-link-name">组件</span>
          </a>
          <ul class="nav-main-submenu">
            <li class="nav-main-item">
              <a class="nav-main-link" href>
                <span class="nav-main-link-name">导航菜单</span>
                <span class="nav-main-link-badge badge badge-pill badge-success">new</span>
              </a>
            </li>
          </ul>
        </li>
      </ul>
    </div>
  </nav>
</template>