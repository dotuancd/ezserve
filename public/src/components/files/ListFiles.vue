<template>
    <div class="container">
        <h1>Files</h1>
        <table class="table">
            <thead>
            <tr>
                <th scope="col">#</th>
                <th scope="col">Name</th>
                <th scope="col">Shareable link</th>
                <th scope="col">Creation</th>
            </tr>
            </thead>
            <tbody>
            <tr v-for="file in files">
                <td scope="row">{{file.id}}</td>
                <td>{{file.name}}</td>
                <td>{{file.public_url}}</td>
                <td>{{file.created_at}}</td>
            </tr>
            </tbody>
        </table>
        <pagination v-bind="{pagination, loadPage}"></pagination>
    </div>
</template>

<script>
    import {mapState} from "vuex";
    import Pagination from "../Pagination";

    export default {
        name: "ListFiles",
        components: {Pagination},
        computed: mapState({
            files: state => state.files.files.items,
            pagination: state => state.files.files
        }),
        created () {
            this.$store.dispatch("files/loadPage", 1)
        },
        methods: {
            nextPage() {

            },
            loadPage(page) {
                this.$store.dispatch("files/loadPage", page)
            }
        }
    }
</script>

<style scoped>

</style>