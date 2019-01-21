<template>
    <div class="container">
        <h1>Files</h1>
        <table class="table">
            <thead>
            <tr>
                <!--<th scope="col">#</th>-->
                <th scope="col">Name</th>
                <th scope="col">Shareable link</th>
                <th scope="col">Creation</th>
            </tr>
            </thead>
            <tbody>
            <tr v-for="file in files" v-bind:key="file.id">
                <!--<td scope="row">{{file.id}}</td>-->
                <td><a href="#">{{file.name}}</a></td>
                <td>{{file.public_url}}</td>
                <td>{{humanDate(file.created_at)}}</td>
            </tr>
            </tbody>
        </table>
        <pagination v-bind="{pagination, loadPage, getLink}"></pagination>
    </div>
</template>

<script>
    import {mapState} from "vuex"
    import Pagination from "../Pagination"
    import date from '../../supports/date'

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
            ...date,
            nextPage() {

            },
            loadPage(page) {
                this.$store.dispatch("files/loadPage", page)
            },
            getLink(page) {
                return "/#/files?page=" + page
            }
        }
    }
</script>

<style scoped>

</style>