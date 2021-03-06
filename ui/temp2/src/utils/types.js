export const PropTypes = {
  Boolean: {
    type: Boolean,
    default: false,
    def(def) {
      return {
        type: Boolean,
        default: def
      }
    }
  },
  String: {
    type: String,
    default: null,
    def(def) {
      return {
        type: String,
        default: def
      }
    }
  },
  Array: {
    type: Array,
    default() {
      return null
    },
    def(def) {
      return {
        type: Array,
        default: def
      }
    }
  },
  Object: {
    type: Object,
    default () {
      return null
    },
    def (def) {
      return {
        type: Object,
        default: def
      }
    }
  },
  Number: {
    type: Number,
    default () {
      return 0
    },
    def (def) {
      return {
        type: Number,
        default: def
      }
    }
  }
}