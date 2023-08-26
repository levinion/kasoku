<template>
    <v-text-field variant="solo" density="compact" v-model="search.input" @update:model-value="search.find"
        class="textfield">
        <template v-slot:prepend-inner>
            <v-icon icon="mdi-arrow-left" @click="search.backToLastDir"></v-icon>
        </template>
    </v-text-field>
</template>

<script lang="ts" setup>
import { onMounted } from 'vue';
import { useSearchStore } from '../store/search';
const search = useSearchStore()
onMounted(() => {
    search.initInput()
})
</script>

<style scoped>
.textfield {
    height: 4vh;
    margin-bottom: 2vh;
}
</style>