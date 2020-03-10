<template>
  <div>
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
        <t-block title="Title" subtitle="subtitle" border >
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
        <t-block title="Title" subtitle="subtitle" border rounded >
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
          <template slot="footer">
            Footer content..
          </template>
        </t-block>
      </b-col>
      <b-col cols="6">
        <t-block title="Title" subtitle="subtitle" theme rounded>
          <p>With header background..</p>
          <template slot="footer">
            Footer content..
          </template>
        </t-block>
      </b-col>
      <b-col cols="6">
        <t-block title="Title" subtitle="subtitle" border rounded >
          <p>Bordered block..</p>
          <template slot="footer">
            Footer content..
          </template>
        </t-block>
      </b-col>
      <b-col cols="6">
        <t-block title="Title" subtitle="subtitle" border theme rounded>
          <p>Bordered block with header background..</p>
          <template slot="footer">
            Footer content..
          </template>
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
              <b-button size="sm" variant="light"><t-icon icon="size-fullscreen" pack="si" /></b-button>
              <b-button size="sm" variant="light"><t-icon icon="pin" pack="si" /></b-button>
              <b-button size="sm" variant="light"><t-icon icon="refresh" pack="si" /></b-button>
              <b-button size="sm" variant="light"><t-icon icon="arrow-up" pack="si" /></b-button>
              <b-button size="sm" variant="light"><t-icon icon="close" pack="si" /></b-button>
            </b-button-group>
          </template>
        </t-block>
      </b-col>
      <b-col cols="6">
        <t-block title="Title" subtitle="subtitle" theme border titleRight>
          <p>Potenti elit lectus augue eget iaculis vitae etiam, ullamcorper etiam bibendum ad feugiat magna accumsan dolor, nibh molestie cras hac ac ad massa, fusce ante convallis ante urna molestie vulputate bibendum tempus ante justo arcu erat accumsan adipiscing risus, libero condimentum venenatis sit nisl nisi ultricies sed, fames aliquet consectetur consequat nostra molestie neque nullam scelerisque neque commodo turpis quisque etiam egestas vulputate massa, curabitur tellus massa venenatis congue dolor enim integer luctus, nisi suscipit gravida fames quis vulputate nisi viverra luctus id leo dictum lorem, inceptos nibh orci.</p>
          <template slot="options">
            <b-button-group>
              <b-button size="sm" variant="light"><t-icon icon="size-fullscreen" pack="si" /></b-button>
              <b-button size="sm" variant="light"><t-icon icon="pin" pack="si" /></b-button>
              <b-button size="sm" variant="light"><t-icon icon="refresh" pack="si" /></b-button>
              <b-button size="sm" variant="light"><t-icon icon="arrow-up" pack="si" /></b-button>
              <b-button size="sm" variant="light"><t-icon icon="close" pack="si" /></b-button>
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
              <b-badge variant="warning"><t-icon icon="exclamation-triangle" /></b-badge>
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
              <b-badge variant="warning"><t-icon icon="exclamation-triangle" /></b-badge>
            </t-block-option-item>
          </template>
        </t-block>
      </b-col>
    </b-row>
    <h2 class="content-heading">Themes</h2>
    <b-row>
      <b-col cols="6" v-for="theme of themes" :key="theme" >
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
      themes: ['primary', 'primary-light', 'primary-dark', 'primary-darker', 'gd-dusk', 'gd-fruit', 'gd-aqua', 'gd-sublime', 'gd-sea', 'gd-leaf', 'gd-lake', 'gd-sun', 'default', 'default-light', 'default-dark', 'default-darker', 'xwork', 'xwork-light', 'xwork-dark', 'xwork-darker', 'xmodern', 'xmodern-light', 'xmodern-dark', 'xmodern-darker', 'xeco', 'xeco-light', 'xeco-dark', 'xeco-darker', 'xsmooth', 'xsmooth-light', 'xsmooth-dark', 'xsmooth-darker', 'xinspire', 'xinspire-light', 'xinspire-dark', 'xinspire-darker', 'success', 'info', 'warning', 'danger', 'gray', 'muted', 'gray-darker', 'black']
    }
  }
}
</script>