<template>
  <div class="t-publish-option-tags">
    <span class="tags-show">
      <TIcon icon="tag" pack="si" class="mr-1" />
      <b-badge v-for="(tag, index) of tags" :key="tag" variant="dark">{{tag}} <span @click="deleteTags(index)">&times;</span></b-badge>
    </span>
    <span class="tags-input">
      <b-form-input placeholder="输入标签，按回车键确认" @blur="inputBlur" v-model="inputTag" class="form-control-alt" @keyup="tagKeyHandler" @compositionend="hint"></b-form-input>
      <div class="note-menu" v-if="showHint">
        <ul>
          <li v-for="(hint, index) of hintList" :key="hint" @click="clickHintItem(hint)" :class="index == hintSelect ? 'active' : ''">{{hint}}</li>
        </ul>
      </div>
    </span>
  </div>
</template>
<script>
import { searchTag } from '@/api/tag'
export default {
  model: {
    prop: 'tags',
    event: 'change.tags'
  },
  props: {
    tags: {
      type: Array,
      default() {
        return []
      }
    }
  },
  data() {
    return {
      inputTag: '',
      showHint: false,
      hintList: [],
      hintSelect: 0,
      typing: null
    }
  },
  methods: {
    tagKeyHandler(e) {
      const tag = this.inputTag.trim()
      if (e.keyCode === 40 || e.keyCode === 38) {
        this.chooseHint(e.keyCode === 40 ? 'down' : 'up')
        return
      }
      if (e.keyCode === 13 && tag.length > 0) {
        if (this.tags.includes(tag)) {
          return
        }
        this.tags.push(tag)
        this.inputTag = ''
        return
      }
      this.hint(tag)
    },
    deleteTags(index) {
      this.tags.splice(index,1)
    },
    hint(text) {
      if (this.typing) {
        clearTimeout(this.typing)
      }
      if (!text) {
        this.showHint = false
        return 
      }
      this.typing = setTimeout(() => {
        this.getHintData(text)
      }, 500)
    },
    getHintData(text){
      searchTag(text).then(res => {
        if (!res.result || res.result.length === 0) {
          return 
        }
        this.hintList = res.result
        this.showHint = true
        this.hintSelect = 0
      })
    },
    chooseHint(arrow) {
      if (arrow == 'down') {
        if (this.hintSelect === this.hintList.length - 1) {
          this.hintSelect = 0
        } else {
          this.hintSelect ++
        }
      } else {
        if (this.hintSelect === 0) {
          this.hintSelect = this.hintList.length - 1
        } else {
          this.hintSelect --
        }
      }
      this.inputTag = this.hintList[this.hintSelect]
    },
    clickHintItem(text) {
      this.tags.push(text)
      this.inputTag = ''
    },
    inputBlur() {
      setTimeout(() => {
        this.showHint = false
      }, 200)
    },
    getValue() {
      return this.tags
    }
  }
}
</script>