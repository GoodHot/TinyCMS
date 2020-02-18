<template>
  <div class="t-publish-option">
    <div class="options">
      <TButton class="mr-1" type="light" icon="settings" iconPack="si" @click="$emit('showAdvance', $event)">高级设置</TButton>
      <span class="tags-show">
        <TIcon icon="tag" pack="si" />
        <TBadge type="dark">java</TBadge>
        <TBadge type="dark">java</TBadge>
        <TBadge type="dark">java</TBadge>
      </span>
      <TInput theme="alt" class="tags-input" placeholder="输入标签，按回车键确认" />
    </div>
    <div>
      <TButton class="mr-1" type="success" icon="paper-plane" iconPack="far">立即发布</TButton>
      <TButton type="light" icon="save" iconPack="far">保存为草稿</TButton>
    </div>
  </div>
</template>