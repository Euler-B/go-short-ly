<script>
    import Item from "./Item.svelte";
    import { onMount } from "svelte"
    import { each } from "svelte/internal";

    let golies = []
    onMount(async () => {
        const res = await fetch("http://localhost:3000/goly")
        golies = await res.json()

    })
</script>

{#each golies as goly}
<Item goly={goly}/>
{/each}
