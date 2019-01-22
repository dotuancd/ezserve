<template>
    <div class="container">
        <form action="" class="form" @submit.prevent="save">
            <h1>New file</h1>
            <div class="form-group">
                <input id="title" class="form-control" type="text" v-model="file.title" placeholder="filename with extension">
            </div>
            <div class="form-group">
                <textarea rows="10" id="content" class="form-control" v-model="file.content"></textarea>
            </div>
            <div class="form-group">
                <button class="btn btn-outline-primary">
                    Create
                </button>
                <button class="btn" @click.prevent="$router.back()">
                    Cancel
                </button>
            </div>
        </form>
    </div>
</template>

<script>
    export default {
        name: "file-create",
        data() {
            return {
                file: {
                    title: "",
                    content: ""
                }
            }
        },
        methods: {
            save () {
                this.$store.dispatch("files/create", this.file, () => {
                    alert('saved')
                })
            }
        }
    }
</script>