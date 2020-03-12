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
    linkTheme: PropTypes.String // 'rotate', 'pop', 'shadow'
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
          [`block-link-${this.linkTheme}`]: this.linkTheme
        }
      ]
    },
    headerClasses() {
      return [
        'block-header',
        {
          'block-header-rtl': this.titleRight,
          [`bg-${this.theme}`]: this.theme && this.theme != 'none',
          'block-header-default': !this.theme
        }
      ]
    },
    renderBody () {
      return (
        <div>
          <div class={this.headerClasses}>
            <h3 class="block-title" v-if={this.title || this.subtitle}>
              {this.title}
              <small v-if={this.subtitle}>{this.subtitle}</small>
            </h3>
            <div class="block-options" v-if={this.$slots.options}>
              <slot name="options"></slot>
            </div>
          </div>
          <div class="block-content">
            <slot></slot>
          </div>
          <div v-if={this.$slots.footer} class="block-content block-content-full block-content-sm bg-body-light font-size-sm">
            <slot name="footer"></slot>
          </div>
        </div>
      )
    }
  },
  render () {
    let title = null
    if (this.title || this.subtitle) {
      let subtitle = null
      if (this.subtitle) {
        subtitle = (
          <small> {this.subtitle} </small>
        )
      }
      title = (
        <h3 class="block-title">
          {this.title} 
          {subtitle}
        </h3>
      )
    }
    let options = null
    if (this.$slots.options) {
      options = (
        <div class="block-options" v-if={this.$slots.options}>
          {this.$slots.options}
        </div>
      )
    }
    let footer = null
    if (this.$slots.footer) {
      footer = (
        <div class="block-content block-content-full block-content-sm bg-body-light font-size-sm">
          {this.$slots.footer}
        </div>
      )
    }
    const body = (
      <div>
        <div class={this.headerClasses}>
          {title}
          {options}
        </div>
        <div class="block-content">
          {this.$slots.default}
        </div>
        {footer}
      </div>
    )
    console.log(this.href)
    if (!this.href) {
      return (
        <div class={this.classes}>
          {body}
        </div>
      )
    }
    return (
      <a class={this.classes} href={this.href}>
        {body}
      </a>
    )
  }
}

Block.install = function (Vue) {
  Vue.component(Block.name, Block)
}

export default Block