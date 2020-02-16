<template>
  <TBlock title="Popover" subtitle="example">
    <TPopover title="example" content="example">
      example
    </TPopover>
    <TPopover type="html">
      <template slot="html">
        <TButtonGroup size="sm" type="light" v-slot:default="{size, type, theme}" >
          <TButton :size="size" :type="type" :theme="theme" icon="check-circle" iconPack="far" iconType="primary">全选</TButton>
          <TButton :size="size" :type="type" :theme="theme" icon="circle" iconPack="far" iconType="primary" >反选</TButton>
        </TButtonGroup>
      </template>
      html example
    </TPopover>
  </TBlock>
</template>