// Class to manage the user's "session" on frontend

export class UserState {
    // 'Runes' to make UI update instantly when data changes
    isLoggedIn = $state(false);
    name = $state("");
    email = $state("");
    id = $state<number | null>(null);

    // *****This function will eventually calls Backend Go API
    // For now, mock a login to test UI
    setSession(userData: { id: number, name: string, email: string }) {
        this.isLoggedIn = true;
        this.id = userData.id;
        this.name = userData.name;
        this.email = userData.email;
    }

    clearSession() {
        this.isLoggedIn = false;
        this.id = null;
        this.name = "";
        this.email = "";
    }
}

// export one single instance so every component shares same data
export const userState = new UserState();