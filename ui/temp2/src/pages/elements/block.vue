<template>
  <div>
    <h2 class="content-heading">Square API</h2>
    <b-row>
      <b-col>
        <t-block title="API" subtitle="Use API" border theme rounded>
          <b-row class="push">
            <b-col>
              <div class="table-responsive">
                <table class="table table-bordered table-vcenter mb-0">
                  <thead class="thead-light">
                    <tr>
                      <th style="width: 25%;">Live Test</th>
                      <th style="min-width: 160px;">Mode</th>
                    </tr>
                  </thead>
                  <tbody>
                    <tr>
                      <td>
                        <b-button variant="light" size="sm" block @click="toggleContent">Toggle Content</b-button>
                      </td>
                      <td>
                        <code>content_toggle</code>
                      </td>
                    </tr>
                    <tr>
                      <td>
                        <b-button variant="light" size="sm" block @click="hideContent">Hide Content</b-button>
                      </td>
                      <td>
                        <code>content_hide</code>
                      </td>
                    </tr>
                    <tr>
                      <td>
                        <b-button variant="light" size="sm" block @click="showContent">Show Content</b-button>
                      </td>
                      <td>
                        <code>content_show</code>
                      </td>
                    </tr>
                    <tr>
                      <td>
                        <b-button variant="light" size="sm" block @click="toggleLoading">Toggle State</b-button>
                      </td>
                      <td>
                        <code>state_toggle</code>
                      </td>
                    </tr>
                    <tr>
                      <td>
                        <b-button variant="light" size="sm" block @click="hideLoading">State Loading</b-button>
                      </td>
                      <td>
                        <code>state_loading</code>
                      </td>
                    </tr>
                    <tr>
                      <td>
                        <b-button variant="light" size="sm" block @click="showLoading">State Normal</b-button>
                      </td>
                      <td>
                        <code>state_normal</code>
                      </td>
                    </tr>
                    <tr>
                      <td>
                        <b-button variant="light" size="sm" block @click="toggleFullscreen">Toggle Fullscreen</b-button>
                      </td>
                      <td>
                        <code>fullscreen_toggle</code>
                      </td>
                    </tr>
                    <tr>
                      <td>
                        <b-button variant="light" size="sm" block @click="fullscreenOn">Fullscreen On</b-button>
                      </td>
                      <td>
                        <code>fullscreen_on</code>
                      </td>
                    </tr>
                    <tr>
                      <td>
                        <b-button variant="light" size="sm" block @click="fullscreenOff">Fullscreen Off</b-button>
                      </td>
                      <td>
                        <code>fullscreen_off</code>
                      </td>
                    </tr>
                    <tr>
                      <td>
                        <b-button variant="light" size="sm" block @click="togglePinned">Toggle Pinned</b-button>
                      </td>
                      <td>
                        <code>pinned_toggle</code>
                      </td>
                    </tr>
                    <tr>
                      <td>
                        <b-button variant="light" size="sm" block @click="pinnedOn">Pinned On</b-button>
                      </td>
                      <td>
                        <code>pinned_on</code>
                      </td>
                    </tr>
                    <tr>
                      <td>
                        <b-button variant="light" size="sm" block @click="pinnedOff">Pinned Off</b-button>
                      </td>
                      <td>
                        <code>pinned_off</code>
                      </td>
                    </tr>
                  </tbody>
                </table>
              </div>
            </b-col>
            <b-col>
              <t-block title="Title" subtitle="subtitle" border theme rounded :pinned="api.pinned" :contentVisible="api.contentVisible" :fullscreen="api.fullscreen" :loading="api.loading">
                <template slot="options">
                  <b-button-group>
                    <b-button size="sm" variant="light" @click="toggleFullscreen">
                      <t-icon icon="size-fullscreen" pack="si" />
                    </b-button>
                    <b-button size="sm" variant="light" @click="togglePinned">
                      <t-icon icon="pin" pack="si" />
                    </b-button>
                    <b-button size="sm" variant="light" @click="toggleLoading">
                      <t-icon icon="refresh" pack="si" />
                    </b-button>
                    <b-button size="sm" variant="light" @click="toggleContent">
                      <t-icon icon="arrow-up" pack="si" />
                    </b-button>
                  </b-button-group>
                </template>
                <p>Potenti elit lectus augue eget iaculis vitae etiam, ullamcorper etiam bibendum ad feugiat magna accumsan dolor, nibh molestie cras hac ac ad massa, fusce ante convallis ante urna molestie vulputate bibendum tempus ante justo arcu erat accumsan adipiscing risus, libero condimentum venenatis sit nisl nisi ultricies sed, fames aliquet consectetur consequat nostra molestie neque nullam scelerisque neque commodo turpis quisque etiam egestas vulputate massa, curabitur tellus massa venenatis congue dolor enim integer luctus, nisi suscipit gravida fames quis vulputate nisi viverra luctus id leo dictum lorem, inceptos nibh orci.</p>
              </t-block>
            </b-col>
          </b-row>
        </t-block>
      </b-col>
    </b-row>
    <h2 class="content-heading">Square Blocks</h2>
    <b-row>
      <b-col cols="6">
        <t-block title="Title" subtitle="subtitle">
          <p>Simple block..</p>
        </t-block>
      </b-col>
      <b-col cols="6">
        <t-block title="Title" subtitle="subtitle" theme>
          <p>With header background..</p>
        </t-block>
      </b-col>
      <b-col cols="6">
        <t-block title="Title" subtitle="subtitle" border>
          <p>Bordered block..</p>
        </t-block>
      </b-col>
      <b-col cols="6">
        <t-block title="Title" subtitle="subtitle" border theme>
          <p>Bordered block with header background..</p>
        </t-block>
      </b-col>
    </b-row>
    <h2 class="content-heading">Rounded Blocks</h2>
    <b-row>
      <b-col cols="6">
        <t-block title="Title" subtitle="subtitle" rounded>
          <p>Simple block..</p>
        </t-block>
      </b-col>
      <b-col cols="6">
        <t-block title="Title" subtitle="subtitle" theme rounded>
          <p>With header background..</p>
        </t-block>
      </b-col>
      <b-col cols="6">
        <t-block title="Title" subtitle="subtitle" border rounded>
          <p>Bordered block..</p>
        </t-block>
      </b-col>
      <b-col cols="6">
        <t-block title="Title" subtitle="subtitle" border theme rounded>
          <p>Bordered block with header background..</p>
        </t-block>
      </b-col>
    </b-row>
    <h2 class="content-heading">Blocks with Footer</h2>
    <b-row>
      <b-col cols="6">
        <t-block title="Title" subtitle="subtitle" rounded>
          <p>Simple block..</p>
          <template slot="footer">Footer content..</template>
        </t-block>
      </b-col>
      <b-col cols="6">
        <t-block title="Title" subtitle="subtitle" theme rounded>
          <p>With header background..</p>
          <template slot="footer">Footer content..</template>
        </t-block>
      </b-col>
      <b-col cols="6">
        <t-block title="Title" subtitle="subtitle" border rounded>
          <p>Bordered block..</p>
          <template slot="footer">Footer content..</template>
        </t-block>
      </b-col>
      <b-col cols="6">
        <t-block title="Title" subtitle="subtitle" border theme rounded>
          <p>Bordered block with header background..</p>
          <template slot="footer">Footer content..</template>
        </t-block>
      </b-col>
    </b-row>
    <h2 class="content-heading">Transparent Blocks</h2>
    <b-row>
      <b-col cols="6">
        <t-block title="Title" subtitle="subtitle" transparent>
          <p>Dolor posuere proin blandit accumsan senectus netus nullam curae, ornare laoreet adipiscing luctus mauris adipiscing pretium eget fermentum, tristique lobortis est ut metus lobortis tortor tincidunt himenaeos habitant quis dictumst proin odio sagittis purus mi, nec taciti vestibulum quis in sit varius lorem sit metus mi.</p>
        </t-block>
      </b-col>
      <b-col cols="6">
        <t-block title="Title" subtitle="subtitle" transparent>
          <p>Dolor posuere proin blandit accumsan senectus netus nullam curae, ornare laoreet adipiscing luctus mauris adipiscing pretium eget fermentum, tristique lobortis est ut metus lobortis tortor tincidunt himenaeos habitant quis dictumst proin odio sagittis purus mi, nec taciti vestibulum quis in sit varius lorem sit metus mi.</p>
        </t-block>
      </b-col>
    </b-row>
    <h2 class="content-heading">Link Blocks</h2>
    <b-row>
      <b-col cols="6">
        <t-block title="Title" subtitle="subtitle" href="#">
          <p>Default opacity hover effect..</p>
        </t-block>
      </b-col>
      <b-col cols="6">
        <t-block title="Title" subtitle="subtitle" href="#" linkTheme="rotate">
          <p>Default opacity hover effect..</p>
        </t-block>
      </b-col>
      <b-col cols="6">
        <t-block title="Title" subtitle="subtitle" href="#" linkTheme="pop">
          <p>Default opacity hover effect..</p>
        </t-block>
      </b-col>
      <b-col cols="6">
        <t-block title="Title" subtitle="subtitle" href="#" linkTheme="shadow">
          <p>Default opacity hover effect..</p>
        </t-block>
      </b-col>
    </b-row>
    <h2 class="content-heading">Block Options</h2>
    <b-row>
      <b-col cols="6">
        <t-block title="Title" subtitle="subtitle" theme border>
          <p>Potenti elit lectus augue eget iaculis vitae etiam, ullamcorper etiam bibendum ad feugiat magna accumsan dolor, nibh molestie cras hac ac ad massa, fusce ante convallis ante urna molestie vulputate bibendum tempus ante justo arcu erat accumsan adipiscing risus, libero condimentum venenatis sit nisl nisi ultricies sed, fames aliquet consectetur consequat nostra molestie neque nullam scelerisque neque commodo turpis quisque etiam egestas vulputate massa, curabitur tellus massa venenatis congue dolor enim integer luctus, nisi suscipit gravida fames quis vulputate nisi viverra luctus id leo dictum lorem, inceptos nibh orci.</p>
          <template slot="options">
            <b-button-group>
              <b-button size="sm" variant="light">
                <t-icon icon="size-fullscreen" pack="si" />
              </b-button>
              <b-button size="sm" variant="light">
                <t-icon icon="pin" pack="si" />
              </b-button>
              <b-button size="sm" variant="light">
                <t-icon icon="refresh" pack="si" />
              </b-button>
              <b-button size="sm" variant="light">
                <t-icon icon="arrow-up" pack="si" />
              </b-button>
              <b-button size="sm" variant="light">
                <t-icon icon="close" pack="si" />
              </b-button>
            </b-button-group>
          </template>
        </t-block>
      </b-col>
      <b-col cols="6">
        <t-block title="Title" subtitle="subtitle" theme border titleRight>
          <p>Potenti elit lectus augue eget iaculis vitae etiam, ullamcorper etiam bibendum ad feugiat magna accumsan dolor, nibh molestie cras hac ac ad massa, fusce ante convallis ante urna molestie vulputate bibendum tempus ante justo arcu erat accumsan adipiscing risus, libero condimentum venenatis sit nisl nisi ultricies sed, fames aliquet consectetur consequat nostra molestie neque nullam scelerisque neque commodo turpis quisque etiam egestas vulputate massa, curabitur tellus massa venenatis congue dolor enim integer luctus, nisi suscipit gravida fames quis vulputate nisi viverra luctus id leo dictum lorem, inceptos nibh orci.</p>
          <template slot="options">
            <b-button-group>
              <b-button size="sm" variant="light">
                <t-icon icon="size-fullscreen" pack="si" />
              </b-button>
              <b-button size="sm" variant="light">
                <t-icon icon="pin" pack="si" />
              </b-button>
              <b-button size="sm" variant="light">
                <t-icon icon="refresh" pack="si" />
              </b-button>
              <b-button size="sm" variant="light">
                <t-icon icon="arrow-up" pack="si" />
              </b-button>
              <b-button size="sm" variant="light">
                <t-icon icon="close" pack="si" />
              </b-button>
            </b-button-group>
          </template>
        </t-block>
      </b-col>
      <b-col cols="6">
        <t-block title="Title" subtitle="subtitle" theme border>
          <p>Potenti elit lectus augue eget iaculis vitae etiam, ullamcorper etiam bibendum ad feugiat magna accumsan dolor, nibh molestie cras hac ac ad massa, fusce ante convallis ante urna molestie vulputate bibendum tempus ante justo arcu erat accumsan adipiscing risus, libero condimentum venenatis sit nisl nisi ultricies sed, fames aliquet consectetur consequat nostra molestie neque nullam scelerisque neque commodo turpis quisque etiam egestas vulputate massa, curabitur tellus massa venenatis congue dolor enim integer luctus, nisi suscipit gravida fames quis vulputate nisi viverra luctus id leo dictum lorem, inceptos nibh orci.</p>
          <template slot="options">
            <b-button size="sm" variant="outline-secondary" class="mr-1">OK</b-button>
            <b-button size="sm" variant="outline-primary">Cancle</b-button>
          </template>
        </t-block>
      </b-col>
      <b-col cols="6">
        <t-block title="Title" subtitle="subtitle" theme border titleRight>
          <p>Potenti elit lectus augue eget iaculis vitae etiam, ullamcorper etiam bibendum ad feugiat magna accumsan dolor, nibh molestie cras hac ac ad massa, fusce ante convallis ante urna molestie vulputate bibendum tempus ante justo arcu erat accumsan adipiscing risus, libero condimentum venenatis sit nisl nisi ultricies sed, fames aliquet consectetur consequat nostra molestie neque nullam scelerisque neque commodo turpis quisque etiam egestas vulputate massa, curabitur tellus massa venenatis congue dolor enim integer luctus, nisi suscipit gravida fames quis vulputate nisi viverra luctus id leo dictum lorem, inceptos nibh orci.</p>
          <template slot="options">
            <b-button size="sm" variant="outline-secondary" class="mr-1">OK</b-button>
            <b-button size="sm" variant="outline-primary">Cancle</b-button>
          </template>
        </t-block>
      </b-col>
      <b-col cols="6">
        <t-block title="Title" subtitle="subtitle" theme border>
          <p>Potenti elit lectus augue eget iaculis vitae etiam, ullamcorper etiam bibendum ad feugiat magna accumsan dolor, nibh molestie cras hac ac ad massa, fusce ante convallis ante urna molestie vulputate bibendum tempus ante justo arcu erat accumsan adipiscing risus, libero condimentum venenatis sit nisl nisi ultricies sed, fames aliquet consectetur consequat nostra molestie neque nullam scelerisque neque commodo turpis quisque etiam egestas vulputate massa, curabitur tellus massa venenatis congue dolor enim integer luctus, nisi suscipit gravida fames quis vulputate nisi viverra luctus id leo dictum lorem, inceptos nibh orci.</p>
          <template slot="options">
            <t-block-option-item class="text-danger">Special!</t-block-option-item>
          </template>
        </t-block>
      </b-col>
      <b-col cols="6">
        <t-block title="Title" subtitle="subtitle" theme border titleRight>
          <p>Potenti elit lectus augue eget iaculis vitae etiam, ullamcorper etiam bibendum ad feugiat magna accumsan dolor, nibh molestie cras hac ac ad massa, fusce ante convallis ante urna molestie vulputate bibendum tempus ante justo arcu erat accumsan adipiscing risus, libero condimentum venenatis sit nisl nisi ultricies sed, fames aliquet consectetur consequat nostra molestie neque nullam scelerisque neque commodo turpis quisque etiam egestas vulputate massa, curabitur tellus massa venenatis congue dolor enim integer luctus, nisi suscipit gravida fames quis vulputate nisi viverra luctus id leo dictum lorem, inceptos nibh orci.</p>
          <template slot="options">
            <t-block-option-item class="text-success">Looking Good!</t-block-option-item>
          </template>
        </t-block>
      </b-col>
      <b-col cols="6">
        <t-block title="Title" subtitle="subtitle" theme border>
          <p>Potenti elit lectus augue eget iaculis vitae etiam, ullamcorper etiam bibendum ad feugiat magna accumsan dolor, nibh molestie cras hac ac ad massa, fusce ante convallis ante urna molestie vulputate bibendum tempus ante justo arcu erat accumsan adipiscing risus, libero condimentum venenatis sit nisl nisi ultricies sed, fames aliquet consectetur consequat nostra molestie neque nullam scelerisque neque commodo turpis quisque etiam egestas vulputate massa, curabitur tellus massa venenatis congue dolor enim integer luctus, nisi suscipit gravida fames quis vulputate nisi viverra luctus id leo dictum lorem, inceptos nibh orci.</p>
          <template slot="options">
            <t-block-option-item>
              <b-alert show variant="warning" class="py-2 mb-0">This is an alert!</b-alert>
            </t-block-option-item>
          </template>
        </t-block>
      </b-col>
      <b-col cols="6">
        <t-block title="Title" subtitle="subtitle" theme border titleRight>
          <p>Potenti elit lectus augue eget iaculis vitae etiam, ullamcorper etiam bibendum ad feugiat magna accumsan dolor, nibh molestie cras hac ac ad massa, fusce ante convallis ante urna molestie vulputate bibendum tempus ante justo arcu erat accumsan adipiscing risus, libero condimentum venenatis sit nisl nisi ultricies sed, fames aliquet consectetur consequat nostra molestie neque nullam scelerisque neque commodo turpis quisque etiam egestas vulputate massa, curabitur tellus massa venenatis congue dolor enim integer luctus, nisi suscipit gravida fames quis vulputate nisi viverra luctus id leo dictum lorem, inceptos nibh orci.</p>
          <template slot="options">
            <t-block-option-item>
              <b-alert show variant="warning" class="py-2 mb-0">This is an alert!</b-alert>
            </t-block-option-item>
          </template>
        </t-block>
      </b-col>
      <b-col cols="6">
        <t-block title="Title" subtitle="subtitle" theme border>
          <p>Potenti elit lectus augue eget iaculis vitae etiam, ullamcorper etiam bibendum ad feugiat magna accumsan dolor, nibh molestie cras hac ac ad massa, fusce ante convallis ante urna molestie vulputate bibendum tempus ante justo arcu erat accumsan adipiscing risus, libero condimentum venenatis sit nisl nisi ultricies sed, fames aliquet consectetur consequat nostra molestie neque nullam scelerisque neque commodo turpis quisque etiam egestas vulputate massa, curabitur tellus massa venenatis congue dolor enim integer luctus, nisi suscipit gravida fames quis vulputate nisi viverra luctus id leo dictum lorem, inceptos nibh orci.</p>
          <template slot="options">
            <t-block-option-item>
              <b-badge variant="primary" pill>15</b-badge>
            </t-block-option-item>
            <t-block-option-item>
              <b-badge variant="success">yes</b-badge>
            </t-block-option-item>
            <t-block-option-item>
              <b-badge variant="warning">
                <t-icon icon="exclamation-triangle" />
              </b-badge>
            </t-block-option-item>
          </template>
        </t-block>
      </b-col>
      <b-col cols="6">
        <t-block title="Title" subtitle="subtitle" theme border titleRight>
          <p>Potenti elit lectus augue eget iaculis vitae etiam, ullamcorper etiam bibendum ad feugiat magna accumsan dolor, nibh molestie cras hac ac ad massa, fusce ante convallis ante urna molestie vulputate bibendum tempus ante justo arcu erat accumsan adipiscing risus, libero condimentum venenatis sit nisl nisi ultricies sed, fames aliquet consectetur consequat nostra molestie neque nullam scelerisque neque commodo turpis quisque etiam egestas vulputate massa, curabitur tellus massa venenatis congue dolor enim integer luctus, nisi suscipit gravida fames quis vulputate nisi viverra luctus id leo dictum lorem, inceptos nibh orci.</p>
          <template slot="options">
            <t-block-option-item>
              <b-badge variant="primary" pill>15</b-badge>
            </t-block-option-item>
            <t-block-option-item>
              <b-badge variant="success">yes</b-badge>
            </t-block-option-item>
            <t-block-option-item>
              <b-badge variant="warning">
                <t-icon icon="exclamation-triangle" />
              </b-badge>
            </t-block-option-item>
          </template>
        </t-block>
      </b-col>
    </b-row>
    <h2 class="content-heading">Themes</h2>
    <b-row>
      <b-col cols="6" v-for="theme of themes" :key="theme">
        <t-block title="Theme" :subtitle="theme" :theme="theme">
          <p>Theme is {{theme}}</p>
        </t-block>
      </b-col>
    </b-row>
  </div>
</template>
<script>
export default {
  data() {
    return {
      title: 'Block',
      description: 'Block使用方式',
      api: {
        contentVisible: true,
        loading: false,
        fullscreen: false,
        pinned: false
      },
      themes: [
        "primary",
        "primary-light",
        "primary-dark",
        "primary-darker",
        "gd-dusk",
        "gd-fruit",
        "gd-aqua",
        "gd-sublime",
        "gd-sea",
        "gd-leaf",
        "gd-lake",
        "gd-sun",
        "default",
        "default-light",
        "default-dark",
        "default-darker",
        "xwork",
        "xwork-light",
        "xwork-dark",
        "xwork-darker",
        "xmodern",
        "xmodern-light",
        "xmodern-dark",
        "xmodern-darker",
        "xeco",
        "xeco-light",
        "xeco-dark",
        "xeco-darker",
        "xsmooth",
        "xsmooth-light",
        "xsmooth-dark",
        "xsmooth-darker",
        "xinspire",
        "xinspire-light",
        "xinspire-dark",
        "xinspire-darker",
        "success",
        "info",
        "warning",
        "danger",
        "gray",
        "muted",
        "gray-darker",
        "black"
      ]
    }
  },
  methods: {
    toggleContent () {
      this.api.contentVisible = !this.api.contentVisible
    },
    hideContent () {
      this.api.contentVisible = false
    },
    showContent () {
      this.api.contentVisible = true
    },
    toggleLoading () {
      this.api.loading = !this.api.loading
    },
    hideLoading () {
      this.api.loading = true
    },
    showLoading () {
      this.api.loading = false
    },
    toggleFullscreen () {
      this.api.fullscreen = !this.api.fullscreen
    },
    fullscreenOn () {
      this.api.fullscreen = true
    },
    fullscreenOff () {
      this.api.fullscreen = false
    },
    togglePinned () {
      this.api.pinned = !this.api.pinned
    },
    pinnedOn () {
      this.api.pinned = true
    },
    pinnedOff () {
      this.api.pinned = false
    }
  }
};
</script>