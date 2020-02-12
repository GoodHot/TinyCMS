<template>
  <div :class="classTheme()">
    <div class="d-lg-none">
      <button
        type="button"
        class="btn btn-block btn-light d-flex justify-content-between align-items-center"
        data-toggle="class-toggle"
        data-target="#horizontal-navigation-hover-normal"
        data-class="d-none"
      >
        Menu - Hover Normal
        <i class="fa fa-bars"></i>
      </button>
    </div>
    <div id="horizontal-navigation-hover-normal" class="d-none d-lg-block mt-2 mt-lg-0">
      <ul class="nav-main nav-main-horizontal nav-main-hover">
        <li class="nav-main-item">
          <a class="nav-main-link active" href="be_ui_navigation_horizontal.html">
            <i class="nav-main-link-icon fa fa-home"></i>
            <span class="nav-main-link-name">Home</span>
            <span class="nav-main-link-badge badge badge-pill badge-primary">5</span>
          </a>
        </li>
        <li class="nav-main-heading">Manage</li>
        <li class="nav-main-item">
          <a
            class="nav-main-link nav-main-link-submenu"
            data-toggle="submenu"
            aria-haspopup="true"
            aria-expanded="false"
            href="#"
          >
            <i class="nav-main-link-icon fa fa-briefcase"></i>
            <span class="nav-main-link-name">Products</span>
          </a>
          <ul class="nav-main-submenu">
            <li class="nav-main-item">
              <a
                class="nav-main-link nav-main-link-submenu"
                data-toggle="submenu"
                aria-haspopup="true"
                aria-expanded="false"
                href="#"
              >
                <i class="nav-main-link-icon fab fa-html5"></i>
                <span class="nav-main-link-name">HTML Templates</span>
              </a>
              <ul class="nav-main-submenu">
                <li class="nav-main-item">
                  <a class="nav-main-link" href="javascript:void(0)">
                    <i class="nav-main-link-icon fa fa-pencil-alt"></i>
                    <span class="nav-main-link-name">Description</span>
                  </a>
                </li>
                <li class="nav-main-item">
                  <a class="nav-main-link" href="javascript:void(0)">
                    <i class="nav-main-link-icon fa fa-chart-line"></i>
                    <span class="nav-main-link-name">Statistics</span>
                  </a>
                </li>
                <li class="nav-main-item">
                  <a class="nav-main-link" href="javascript:void(0)">
                    <i class="nav-main-link-icon fa fa-chart-area"></i>
                    <span class="nav-main-link-name">Sales</span>
                    <span class="nav-main-link-badge badge badge-pill badge-primary">320</span>
                  </a>
                </li>
                <li class="nav-main-item">
                  <a class="nav-main-link" href="javascript:void(0)">
                    <i class="nav-main-link-icon far fa-images"></i>
                    <span class="nav-main-link-name">Media</span>
                    <span class="nav-main-link-badge badge badge-pill badge-primary">18</span>
                  </a>
                </li>
                <li class="nav-main-item">
                  <a class="nav-main-link" href="javascript:void(0)">
                    <i class="nav-main-link-icon far fa-file-alt"></i>
                    <span class="nav-main-link-name">Files</span>
                    <span class="nav-main-link-badge badge badge-pill badge-primary">32</span>
                  </a>
                </li>
              </ul>
            </li>
            <li class="nav-main-item">
              <a
                class="nav-main-link nav-main-link-submenu"
                data-toggle="submenu"
                aria-haspopup="true"
                aria-expanded="false"
                href="#"
              >
                <i class="nav-main-link-icon fab fa-wordpress"></i>
                <span class="nav-main-link-name">WordPress Themes</span>
              </a>
              <ul class="nav-main-submenu">
                <li class="nav-main-item">
                  <a class="nav-main-link" href="javascript:void(0)">
                    <i class="nav-main-link-icon fa fa-pencil-alt"></i>
                    <span class="nav-main-link-name">Description</span>
                  </a>
                </li>
                <li class="nav-main-item">
                  <a class="nav-main-link" href="javascript:void(0)">
                    <i class="nav-main-link-icon fa fa-chart-line"></i>
                    <span class="nav-main-link-name">Statistics</span>
                  </a>
                </li>
                <li class="nav-main-item">
                  <a class="nav-main-link" href="javascript:void(0)">
                    <i class="nav-main-link-icon fa fa-chart-area"></i>
                    <span class="nav-main-link-name">Sales</span>
                    <span class="nav-main-link-badge badge badge-pill badge-primary">680</span>
                  </a>
                </li>
                <li class="nav-main-item">
                  <a class="nav-main-link" href="javascript:void(0)">
                    <i class="nav-main-link-icon far fa-images"></i>
                    <span class="nav-main-link-name">Media</span>
                    <span class="nav-main-link-badge badge badge-pill badge-primary">15</span>
                  </a>
                </li>
                <li class="nav-main-item">
                  <a class="nav-main-link" href="javascript:void(0)">
                    <i class="nav-main-link-icon far fa-file-alt"></i>
                    <span class="nav-main-link-name">Files</span>
                    <span class="nav-main-link-badge badge badge-pill badge-primary">20</span>
                  </a>
                </li>
              </ul>
            </li>
            <li class="nav-main-item">
              <a
                class="nav-main-link nav-main-link-submenu"
                data-toggle="submenu"
                aria-haspopup="true"
                aria-expanded="false"
                href="#"
              >
                <i class="nav-main-link-icon fab fa-medapps"></i>
                <span class="nav-main-link-name">Web Applications</span>
              </a>
              <ul class="nav-main-submenu">
                <li class="nav-main-item">
                  <a class="nav-main-link" href="javascript:void(0)">
                    <i class="nav-main-link-icon fa fa-pencil-alt"></i>
                    <span class="nav-main-link-name">Description</span>
                  </a>
                </li>
                <li class="nav-main-item">
                  <a class="nav-main-link" href="javascript:void(0)">
                    <i class="nav-main-link-icon fa fa-chart-line"></i>
                    <span class="nav-main-link-name">Statistics</span>
                  </a>
                </li>
                <li class="nav-main-item">
                  <a class="nav-main-link" href="javascript:void(0)">
                    <i class="nav-main-link-icon fa fa-chart-area"></i>
                    <span class="nav-main-link-name">Sales</span>
                    <span class="nav-main-link-badge badge badge-pill badge-primary">158</span>
                  </a>
                </li>
                <li class="nav-main-item">
                  <a class="nav-main-link" href="javascript:void(0)">
                    <i class="nav-main-link-icon far fa-images"></i>
                    <span class="nav-main-link-name">Media</span>
                    <span class="nav-main-link-badge badge badge-pill badge-primary">65</span>
                  </a>
                </li>
                <li class="nav-main-item">
                  <a class="nav-main-link" href="javascript:void(0)">
                    <i class="nav-main-link-icon far fa-file-alt"></i>
                    <span class="nav-main-link-name">Files</span>
                    <span class="nav-main-link-badge badge badge-pill badge-primary">78</span>
                  </a>
                </li>
              </ul>
            </li>
            <li class="nav-main-item">
              <a
                class="nav-main-link nav-main-link-submenu"
                data-toggle="submenu"
                aria-haspopup="true"
                aria-expanded="false"
                href="#"
              >
                <i class="nav-main-link-icon fab fa-youtube"></i>
                <span class="nav-main-link-name">Video Templates</span>
              </a>
              <ul class="nav-main-submenu">
                <li class="nav-main-item">
                  <a class="nav-main-link" href="javascript:void(0)">
                    <i class="nav-main-link-icon fa fa-pencil-alt"></i>
                    <span class="nav-main-link-name">Description</span>
                  </a>
                </li>
                <li class="nav-main-item">
                  <a class="nav-main-link" href="javascript:void(0)">
                    <i class="nav-main-link-icon fa fa-chart-line"></i>
                    <span class="nav-main-link-name">Statistics</span>
                  </a>
                </li>
                <li class="nav-main-item">
                  <a class="nav-main-link" href="javascript:void(0)">
                    <i class="nav-main-link-icon fa fa-chart-area"></i>
                    <span class="nav-main-link-name">Sales</span>
                    <span class="nav-main-link-badge badge badge-pill badge-primary">920</span>
                  </a>
                </li>
                <li class="nav-main-item">
                  <a class="nav-main-link" href="javascript:void(0)">
                    <i class="nav-main-link-icon far fa-images"></i>
                    <span class="nav-main-link-name">Media</span>
                    <span class="nav-main-link-badge badge badge-pill badge-primary">7</span>
                  </a>
                </li>
                <li class="nav-main-item">
                  <a class="nav-main-link" href="javascript:void(0)">
                    <i class="nav-main-link-icon far fa-file-alt"></i>
                    <span class="nav-main-link-name">Files</span>
                    <span class="nav-main-link-badge badge badge-pill badge-primary">19</span>
                  </a>
                </li>
              </ul>
            </li>
            <li class="nav-main-item">
              <a class="nav-main-link" href="javascript:void(0)">
                <i class="nav-main-link-icon fa fa-plus"></i>
                <span class="nav-main-link-name">Add Product</span>
              </a>
            </li>
          </ul>
        </li>
        <li class="nav-main-item">
          <a
            class="nav-main-link nav-main-link-submenu"
            data-toggle="submenu"
            aria-haspopup="true"
            aria-expanded="false"
            href="#"
          >
            <i class="nav-main-link-icon fa fa-money-bill-wave"></i>
            <span class="nav-main-link-name">Payments</span>
          </a>
          <ul class="nav-main-submenu">
            <li class="nav-main-item">
              <a class="nav-main-link" href="javascript:void(0)">
                <span class="nav-main-link-name">Scheduled</span>
                <span class="nav-main-link-badge badge badge-pill badge-success">3</span>
              </a>
            </li>
            <li class="nav-main-item">
              <a class="nav-main-link" href="javascript:void(0)">
                <span class="nav-main-link-name">Recurring</span>
              </a>
            </li>
            <li class="nav-main-item">
              <a class="nav-main-link" href="javascript:void(0)">
                <span class="nav-main-link-name">Manage</span>
              </a>
            </li>
            <li class="nav-main-item">
              <a class="nav-main-link" href="javascript:void(0)">
                <i class="nav-main-link-icon fa fa-plus"></i>
                <span class="nav-main-link-name">New Payment</span>
              </a>
            </li>
          </ul>
        </li>
        <li class="nav-main-item">
          <a
            class="nav-main-link nav-main-link-submenu"
            data-toggle="submenu"
            aria-haspopup="true"
            aria-expanded="false"
            href="#"
          >
            <i class="nav-main-link-icon fa fa-globe-americas"></i>
            <span class="nav-main-link-name">Services</span>
          </a>
          <ul class="nav-main-submenu">
            <li class="nav-main-item">
              <a class="nav-main-link" href="javascript:void(0)">
                <span class="nav-main-link-name">Hosting</span>
              </a>
            </li>
            <li class="nav-main-item">
              <a class="nav-main-link" href="javascript:void(0)">
                <span class="nav-main-link-name">Web Design</span>
              </a>
            </li>
            <li class="nav-main-item">
              <a class="nav-main-link" href="javascript:void(0)">
                <span class="nav-main-link-name">Web Development</span>
              </a>
            </li>
            <li class="nav-main-item">
              <a class="nav-main-link" href="javascript:void(0)">
                <span class="nav-main-link-name">Graphic Design</span>
              </a>
            </li>
            <li class="nav-main-item">
              <a class="nav-main-link" href="javascript:void(0)">
                <span class="nav-main-link-name">Legal</span>
              </a>
            </li>
            <li class="nav-main-item">
              <a class="nav-main-link" href="javascript:void(0)">
                <span class="nav-main-link-name">Consulting</span>
              </a>
            </li>
          </ul>
        </li>
        <li class="nav-main-heading">Personal</li>
        <li class="nav-main-item">
          <a
            class="nav-main-link nav-main-link-submenu"
            data-toggle="submenu"
            aria-haspopup="true"
            aria-expanded="false"
            href="#"
          >
            <i class="nav-main-link-icon far fa-user"></i>
            <span class="nav-main-link-name">Profile</span>
          </a>
          <ul class="nav-main-submenu">
            <li class="nav-main-item">
              <a class="nav-main-link" href="javascript:void(0)">
                <span class="nav-main-link-name">Edit</span>
              </a>
            </li>
            <li class="nav-main-item">
              <a class="nav-main-link" href="javascript:void(0)">
                <span class="nav-main-link-name">Settings</span>
              </a>
            </li>
            <li class="nav-main-item">
              <a class="nav-main-link" href="javascript:void(0)">
                <span class="nav-main-link-name">Log out</span>
              </a>
            </li>
          </ul>
        </li>
      </ul>
    </div>
  </div>
</template>
<script>
export default {
  props: {
    theme: {
      type: String,
      default: "light"
    }
  },
  methods: {
    classTheme() {
      if (this.theme === "light") {
        return "bg-white p-3 push"
      } else {
        return "bg-sidebar-dark p-3 push"
      }
    }
  }
}
</script>