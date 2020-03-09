export const PropTypes = {
  Boolean: {
    type: Boolean,
    default: false
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