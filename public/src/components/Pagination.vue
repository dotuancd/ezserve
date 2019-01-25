<template>
    <nav aria-label="Page navigation example">
        <ul class="pagination justify-content-end">
            <li class="page-item" :class="{disabled: !hasPrevious}">
                <a class="page-link" :href="getLink(pagination.page - 1)" @click.prevent="loadPage(pagination.page-1)">Previous</a>
            </li>
            <li class="page-item" v-for="i in pagination.total_page" :class="{active: isCurrent(i)}">
                <a class="page-link" :href="getLink(i)" @click.prevent="loadPage(i)">{{i}}</a>
            </li>
            <li class="page-item" :class="{disabled: !hasNext}">
                <a class="page-link" v-bind:href="getLink(pagination.page + 1)" @click.prevent="loadPage(pagination.page + 1)">Next</a>
            </li>
        </ul>
    </nav>
</template>

<script>
    export default {
        name: "Pagination",
        computed: {
            hasPrevious() {
                return (this.pagination.page !== 1)
            },

            hasNext() {
                return (this.pagination.page !== this.pagination.total_page)
            }
        },
        methods: {
            isCurrent(page) {
                return (this.pagination.page === page)
            },
        },
        props: ["pagination", "loadPage", "getLink"]
    }
</script>

<style scoped>

</style>