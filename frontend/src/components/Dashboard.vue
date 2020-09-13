<template>
    <v-card width="100%" tile>
        <v-card-title>
            <v-text-field v-model="search"
                append-icon="mdi-magnify"
                label="Search"
                single-line
                hide-details
            ></v-text-field>
        </v-card-title>
        <v-data-table
            :headers="headers"
            :items="people"
            :search="search"
        ></v-data-table>
    </v-card>
</template>

<script>
import axios from 'axios';

export default {
    name: 'App',
    components: {
    },
    data() {
        return {
            people: [],
            search: "",
            headers: [
                {text: "First Name", align: "start", value: "first_name"},
                {text: "Last Name", value: "last_name"},
                {text: "Date of Birth", value: "dob"},
                {text: "Email", value: "email"},
                {text: "Phone", value: "phone"}
            ]
        }
    },
    mounted() {
        this.getAllElements()
    },
    methods: {
        getAllElements: function () {
            axios.get(`/app/profiles`)
                .then(response => {
                    this.people = response.data
                })
                .catch(error => {
                    console.log(error)
                })
        }
    }
}
</script>

<style scoped>

</style>