<template>
  <select v-on:change="change" :value="method">
    <option value="periodic">{{ t("settings.uploadingMethod.scheduleUpload") }}</option>
    <option value="onetime">{{ t("settings.uploadingMethod.oneTimeUpload") }}</option>
  </select>
</template>

<script setup lang="ts">
import type { SelectHTMLAttributes } from "vue";
import { useI18n } from "vue-i18n";

const { t } = useI18n();

defineProps<{
  method: UploadingMethod;
}>();

const emit = defineEmits<{
  (e: "update:method", val: string | null): void;
}>();

const change = (event: Event) => {
  emit("update:method", (event.target as SelectHTMLAttributes)?.value);
};
</script>
