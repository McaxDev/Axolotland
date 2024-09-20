<script setup lang="ts"></script>
<template>
  <div class="container-title" style="margin-top: 200px;">Fixed Width Container</div>
  <!-- 固定宽度容器 -->
  <b-container class="bg-dark">
        This is a fixed-width container. The width is fixed and does not stretch to fill the entire
        screen.
  </b-container>

  <div class="container-title">Fluid Container</div>
  <!-- 流体容器 -->
  <b-container fluid class="bg-dark">
    This is a fluid container. The width stretches to fill the entire width of the screen.
  </b-container>
</template>
