<template>
  <section class="app-main w-full bg-slate-100 relative overflow-hidden min-h-screen">
    <router-view v-slot="{ Component }">
      <keep-alive>
        <component :is="Component" />
      </keep-alive>
    </router-view>
  </section>
</template>
