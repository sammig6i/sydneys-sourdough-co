'use client'
import { Button } from '@/components/ui/button';
import {
  Card,
  CardDescription,
  CardFooter,
  CardHeader,
  CardTitle
} from '@/components/ui/card';
import { handleSignInWithGoogle } from '@/lib/auth';
import { useEffect } from 'react';

export default function LoginPage() {
  useEffect(() => {
    window.handleSignInWithGoogle = handleSignInWithGoogle;
  }, []);

  return (
    <div className="min-h-screen flex justify-center items-start md:items-center p-8">
      <Card className="w-full max-w-sm">
        <CardHeader>
          <CardTitle className="text-2xl">Login</CardTitle>
        </CardHeader>
        <CardFooter>
          <Button id="g_id_onload"
            data-client_id="601595052196-uucvqf6jl1rmn4dmjjhbq9s7cghfpprm.apps.googleusercontent.com"
            data-context="signin"
            data-ux_mode="popup"
            data-callback="handleSignInWithGoogle"
            data-auto_prompt="false"
            className="w-full g_id_signin"
            data-type="standard"
            data-shape="rectangular"
            data-theme="outline"
            data-text="signin_with"
            data-size="large"
            data-logo_alignment="left"
            data-use_fedcm_for_prompt="true">
          </Button>
        </CardFooter>
      </Card>
    </div>
  );
}
