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
  }
}