import {PropTypes} from '@/utils/types'

const Block = {
  name: 'block',
  props: {
    title: PropTypes.String,
    subtitle: PropTypes.String,
    border: PropTypes.Boolean,
    rounded: PropTypes.Boolean,
    transparent: PropTypes.Boolean,
    titleRight: PropTypes.Boolean,
    theme: PropTypes.String.def('none'), // 'primary-light', 'primary-dark', 'primary-darker', 'gd-dusk', 'gd-fruit', 'gd-aqua', 'gd-sublime', 'gd-sea', 'gd-leaf', 'gd-lake', 'gd-sun', 'default', 'default-light', 'default-dark', 'default-darker', 'xwork', 'xwork-light', 'xwork-dark', 'xwork-darker', 'xmodern', 'xmodern-light', 'xmodern-dark', 'xmodern-darker', 'xeco', 'xeco-light', 'xeco-dark', 'xeco-darker', 'xsmooth', 'xsmooth-light', 'xsmooth-dark', 'xsmooth-darker', 'xinspire', 'xinspire-light', 'xinspire-dark', 'xinspire-darker', 'success', 'info', 'warning', 'danger', 'gray', 'muted', 'gray-darker', 'black', 
    href: PropTypes.String,
    linkTheme: PropTypes.String, // 'rotate', 'pop', 'shadow'
    contentVisible: PropTypes.Boolean.def(true),
    loading: PropTypes.Boolean,
    fullscreen: PropTypes.Boolean,
    pinned: PropTypes.Boolean
  },
  computed: {
    classes() {
      return [
        'block',
        {
          'block-themed': this.theme && this.theme != 'none',
          'block-bordered': this.border,
          'block-rounded': this.rounded,
          'block-transparent': this.transparent,
          'block-mode-loading': this.loading,
          'block-mode-fullscreen': this.fullscreen,
          'block-mode-pinned': this.pinned,
          [`block-link-${this.linkTheme}`]: this.linkTheme
        }
      ]
    }
  },
  watch: {
  },
  render () {
    const blockContent = (
      <t-block-content
        title={this.title}
        subtitle={this.subtitle}
        border={this.border}
        rounded={this.rounded}
        transparent={this.transparent}
        titleRight={this.titleRight}
        theme={this.theme}
        linkTheme={this.linkTheme}
        contentVisible={this.contentVisible}
        loading={this.loading}
      >
        <template slot="default">
          {this.$slots.default}
        </template>
        <template slot="options">
          {this.$slots.options}
        </template>
        <template slot="footer">
          {this.$slots.footer}
        </template>
      </t-block-content>
    )
    if (!this.href) {
      return (
        <div class={this.classes}>
          {blockContent}
        </div>
      )
    }
    return (
      <a class={this.classes} href={this.href}>
        {blockContent}
      </a>
    )
  }
}

Block.install = function (Vue) {
  Vue.component(Block.name, Block)
}

export default Block