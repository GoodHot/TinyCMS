<template>
  <b-row>
    <b-col cols="3">
      <t-block theme border title="操作" subtitle="API">
        <b-button-group size="sm" class="push">
          <b-button @click="toggleHorizontal">触发方向</b-button>
          <b-button @click="options.horizontal = true; col = 9">水平</b-button>
          <b-button @click="options.horizontal = false; col = 3">垂直</b-button>
        </b-button-group>
        <br />
        <b-button-group size="sm" class="push">
          <b-button @click="options.rounded = !options.rounded">触发圆角</b-button>
          <b-button @click="options.rounded = true">圆角</b-button>
          <b-button @click="options.rounded = false">直角</b-button>
        </b-button-group>
        <br />
        <b-button-group size="sm" class="push">
          <b-button @click="toggleTheme">切换主题</b-button>
          <b-button @click="options.theme = 'light'">light</b-button>
          <b-button @click="options.theme = 'dark'">dark</b-button>
        </b-button-group>
        <br>
        <b-button-group size="sm" class="push">
          <b-button @click="options.layout = 'center'">居中</b-button>
          <b-button @click="options.layout = 'left'">居右</b-button>
          <b-button @click="options.layout = 'full'">铺满</b-button>
        </b-button-group>
        <br>
        <b-button-group size="sm" class="push">
          <b-button @click="options.showEvent = 'hover'">移动</b-button>
          <b-button @click="options.showEvent = 'click'">点击</b-button>
        </b-button-group>
      </t-block>
    </b-col>
    <b-col :cols="col">
      <t-navigation
        :horizontal="options.horizontal"
        :theme="options.theme"
        :rounded="options.rounded"
        :data="navigation"
        :layout="options.layout"
        :showEvent="options.showEvent"
      >
      </t-navigation>
    </b-col>
  </b-row>
</template>
<script>
const navChild = [
  {
    title: "Description",
    icon: "pencil-alt"
  },
  {
    title: "Statistics",
    icon: "chart-pie",
    child: [
      {
        title: "HTML Template",
        icon: "code"
      },
      {
        title: "Description",
        icon: "pencil-alt"
      },
      {
        title: "Statistics",
        icon: "chart-pie",
        child: [
          {
            title: "Description",
            icon: "pencil-alt"
          },
          {
            title: "Statistics",
            icon: "chart-pie"
          },
          {
            title: "Sales",
            icon: "chart-line"
          },
          {
            title: "Description",
            icon: "pencil-alt"
          },
          {
            title: "Statistics",
            icon: "chart-pie"
          },
          {
            title: "Sales",
            icon: "chart-line"
          },
          {
            title: "New Product",
            icon: "plus-circle"
          }
        ]
      },
      {
        title: "Sales",
        icon: "chart-line"
      }
    ]
  },
  {
    title: "Sales",
    icon: "chart-line"
  }
];
export default {
  data() {
    return {
      col: 9,
      options: {
        horizontal: true,
        rounded: false,
        theme: 'light',
        layout: 'left',
        showEvent: 'hover'
      },
      navigation: [
        {
          title: "Overview",
          icon: "rocket",
          badgeText: "3"
        },
        {
          title: "Product",
          icon: "boxes",
          child: navChild
        },
        {
          title: "Payments",
          icon: "money-bill",
          child: navChild
        },
        {
          title: "Services",
          icon: "globe",
          child: navChild
        },
        {
          title: "Profile",
          icon: "user-circle"
        }
      ]
    };
  },
  methods: {
    toggleHorizontal() {
      this.options.horizontal = !this.options.horizontal;
      if (this.options.horizontal) {
        this.col = 9
      } else {
        this.col = 3
      }
    },
    toggleTheme() {
      this.options.theme = this.options.theme === 'dark' ? 'light' : 'dark'
    }
  }
};
</script>