<script>
    import { MessageCircle, CheckCircle, UserCircle } from "@lucide/svelte";

    export let title = "UI/UX Design in the age of AI";
    export let description = "Lorem ipsum dolor sit amet, libre unst consectetur adipiscing elit.";
    export let comments = 11;
    export let tasks = 187;
    export let badgeText = "Important";
</script>

<button class="w-full p-4 bg-white border border-gray-200 rounded-xl shadow-sm hover:shadow-lg transition cursor-pointer">
    <!-- Badge -->
    <div class="text-xs font-semibold text-indigo-600 bg-indigo-100 px-3 py-1 rounded-full w-fit mb-2">
        {badgeText}
    </div>

    <div class="mb-3">
        <h2 class="text-lg font-semibold text-gray-900">{title}</h2>
        <p class="text-sm text-gray-600">{description}</p>
    </div>

    <div class="flex justify-between items-center mt-3 text-sm text-gray-500">
        <!-- Avatars (Using Lucide Icons) -->
        <div class="flex">
            <UserCircle class="w-8 h-8 text-gray-400 border-2 border-white bg-gray-100 rounded-full -mr-2" />
            <UserCircle class="w-8 h-8 text-gray-400 border-2 border-white bg-gray-100 rounded-full" />
        </div>

        <!-- Comments & Tasks -->
        <div class="flex gap-4">
            <div class="flex items-center gap-1">
                <MessageCircle size={16} class="text-gray-500" />
                <span>{comments}</span>
            </div>
            <div class="flex items-center gap-1">
                <CheckCircle size={16} class="text-gray-500" />
                <span>{tasks}</span>
            </div>
        </div>
    </div>
</button>
