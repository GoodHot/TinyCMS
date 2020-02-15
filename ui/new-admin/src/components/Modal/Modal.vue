<template>
  <!--
    Modal
      show: true/false default false
      size: sm,lg,xl
      position: top, center
  -->
  <div class="modal fade" ref="modal">
    <div :class="modalClass">
      <div class="modal-content">
        <div class="block block-themed block-transparent mb-0">
          <div class="block-header bg-primary-dark">
            <h3 class="block-title">{{ title }}</h3>
            <div class="block-options">
              <button
                type="button"
                class="btn-block-option"
                @click="close"
              >
                <i class="fa fa-fw fa-times"></i>
              </button>
            </div>
          </div>
          <div class="block-content font-size-sm">
            <slot></slot>
          </div>
          <div class="block-content block-content-full text-right border-top">
            <TButton class="mr-2" type="light" size="sm" @click="close">Close</TButton>
            <TButton size="sm" @click="okHandler">OK</TButton>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
<script>
export default {
  props: {
    show: {
      type: Boolean,
      default: false
    },
    title: {
      type: String,
      default: 'Title'
    },
    size: {
      type: String,
      default: null
    },
    position: {
      type: String,
      default: null
    }
  },
  data() {
    return {
      visiable: this.show
    }
  },
  computed: {
    modalClass() {
      let cls = 'modal-dialog '
      if (this.size) {
        cls += `modal-${this.size} `
      }
      if (this.position) {
        if (this.position === 'center') {
          cls += 'modal-dialog-centered '
        } else {
          cls += `modal-dialog-${this.position} `
        }
      }
      return cls
    }
  },
  mounted() {
    this.handler()
    const modal = this.$refs.modal;
    const _this = this
    this.$(modal).on('hide.bs.modal', function () {
      _this.close()
    })
  },
  methods: {
    close() {
      this.visiable = false
      this.$emit("onclose")
    },
    okHandler() {
      this.$emit("okHandler")
    },
    handler() {
      const modal = this.$refs.modal;
      if (this.visiable) {
        this.$(modal).modal("show");
      } else {
        this.$(modal).modal("hide");
      }
    }
  },
  watch: {
    show(val) {
      this.visiable = val
      this.handler()
    }
  }
};
</script>