<template>
    <v-text-field variant="solo" density="compact" v-model="search.fuseInput" class="footer" :loading="search.loading"
        ref="fuse" clearable @keydown.enter="search.fuseFind" @keydown.esc="search.clearFuseInput">
        <template v-slot:append-inner>
            <v-icon icon="mdi-send" @click="search.fuseFind"></v-icon>
        </template>
        <template v-slot:prepend-inner>
            <v-icon v-if="search.jumpHiddenItem" icon="mdi-tag-off" @click="search.toggleJumpHiddenItem"></v-icon>
            <v-icon v-else icon="mdi-tag" @click="search.toggleJumpHiddenItem"></v-icon>
        </template>
    </v-text-field>
</template>

<script lang="ts" setup>
import { ref } from 'vue';
import { useSearchStore } from '../store/search';
import { onMounted } from 'vue';
const search = useSearchStore()
const fuse = ref()
onMounted(() => {
    fuse.value.focus()
})
</script>

<style scoped>
.footer {
    height: 3vh;
}
</style>