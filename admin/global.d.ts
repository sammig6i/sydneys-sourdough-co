// global.d.ts
declare global {
  interface Window {
    handleSignInWithGoogle: (response: CredentialResponse) => void;
  }
}

export {};
