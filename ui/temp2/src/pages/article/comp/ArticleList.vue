<template>
  <div class="content">
    <div class="d-md-none push">
      <b-button block variant="hero-primary">
        Article List
      </b-button>
    </div>
    <div class="d-md-block push d-none">
      <b-button block variant="hero-success" class="mb-3"><t-icon icon="plus" class="mr-1" /> 发布新文章</b-button>

      <b-input-group class="mb-3">
        <b-form-input class="border-0" placeholder="文章搜索"></b-form-input>
        <b-input-group-append>
          <b-button class="bg-white px-3" variant="light" size="sm" text="Button"><t-icon size="fw" icon="search" /></b-button>
        </b-input-group-append>
      </b-input-group>

      <div class="d-flex justify-content-between mb-2">
        <b-dropdown text="Sort by" variant="link" size="sm">
          <b-dropdown-item href="#"><t-icon icon="sort-amount-down" size="fw" class="mr-1" />Newest</b-dropdown-item>
          <b-dropdown-item href="#"><t-icon icon="sort-amount-up" size="fw" class="mr-1" />Oldest</b-dropdown-item>
        </b-dropdown>
        <b-dropdown id="dropdown-1" text="Filter by" variant="link" size="sm" class="m-md-2">
          <b-dropdown-item>First Action</b-dropdown-item>
          <b-dropdown-item>Second Action</b-dropdown-item>
        </b-dropdown>
      </div>
      <b-list-group class="font-size-sm">
        <b-list-group-item href="#" class="ribbon ribbon-light">
          <div class="ribbon-box"><t-icon icon="save" pack="far" /></div>
          <p class="font-size-h6 font-w700 mb-0">New Photo PackNew Photo PackNew Photo PackNew Photo PackNew Photo Pack</p>
          <p class="text-muted mb-2">
            Our new photo pack is now available to go live and we couldn’t be more excited..
          </p>
          <small class="text-muted">3days ago create by xww</small>
        </b-list-group-item>
        <b-list-group-item href="#" class="ribbon ribbon-light">
          <div class="ribbon-box"><t-icon icon="save" pack="far" /></div>
          <p class="font-size-h6 font-w700 mb-0">New Photo PackNew Photo PackNew Photo PackNew Photo PackNew Photo Pack</p>
          <p class="text-muted mb-2">
            Our new photo pack is now available to go live and we couldn’t be more excited..
          </p>
          <small class="text-muted">3days ago create by xww</small>
        </b-list-group-item>
        <b-list-group-item href="#" class="ribbon ribbon-light">
          <div class="ribbon-box"><t-icon icon="save" pack="far" /></div>
          <p class="font-size-h6 font-w700 mb-0">New Photo PackNew Photo PackNew Photo PackNew Photo PackNew Photo Pack</p>
          <p class="text-muted mb-2">
            Our new photo pack is now available to go live and we couldn’t be more excited..
          </p>
          <small class="text-muted">3days ago create by xww</small>
        </b-list-group-item>
        <b-list-group-item href="#" class="ribbon ribbon-light">
          <div class="ribbon-box"><t-icon icon="save" pack="far" /></div>
          <p class="font-size-h6 font-w700 mb-0">New Photo PackNew Photo PackNew Photo PackNew Photo PackNew Photo Pack</p>
          <p class="text-muted mb-2">
            Our new photo pack is now available to go live and we couldn’t be more excited..
          </p>
          <small class="text-muted">3days ago create by xww</small>
        </b-list-group-item>
      </b-list-group>
    </div>
  </div>
</template>