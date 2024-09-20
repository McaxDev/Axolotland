<script setup lang="ts">
const props = withDefaults(defineProps<{
	internal?: boolean;
}>(), {
	internal: false,
})
</script>

<template>
	<a><slot></slot>
		<i v-show="!props.internal" class="bi bi-box-arrow-up-right ms-1">
		</i>
	</a>
</template>
