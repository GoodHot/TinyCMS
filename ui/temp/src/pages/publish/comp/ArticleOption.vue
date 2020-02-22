<template>
  <div class="t-publish-option">
    <div class="options">
      <b-dropdown text="分类: 撒大声地" variant="light" class="mr-2 float-left" style="width:250px">
        
      </b-dropdown>
      <b-button class="mr-2" variant="light" @click="$emit('showAdvance', $event)">
        <TIcon icon="settings" pack="si" />
        高级设置
      </b-button>
      <span class="tags-show">
        <TIcon icon="tag" pack="si" class="mr-1" />
        <!-- <b-badge variant="dark">Java</b-badge>
        <b-badge variant="dark">Java</b-badge>
        <b-badge variant="dark">Java</b-badge>
        <b-badge variant="dark">Java</b-badge> -->
      </span>
      <b-form-input placeholder="输入标签，按回车键确认" class="form-control-alt tags-input"></b-form-input>
    </div>
    <div>
      <b-button class="mr-2" variant="success">
        <TIcon icon="paper-plane" pack="far" />
        立即发布
      </b-button>

      <b-button variant="light">
        <TIcon icon="save" pack="far" />
        保存为草稿
      </b-button>
    </div>
  </div>
</template>