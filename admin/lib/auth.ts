import { redirect } from 'next/navigation';
import { createClient } from './client';
import { createServerClient } from '@supabase/ssr';
import { NextResponse, type NextRequest } from 'next/server';
import { revalidatePath } from 'next/cache';
// TODO update Supabase Auth to allow certain users
interface CredentialResponse {
  credential: string;
  select_by: string;
  state: string;
}

export async function handleSignInWithGoogle(response: CredentialResponse) {
  const supabase = await createClient();

  const { data, error } = await supabase.auth.signInWithIdToken({
    provider: 'google',
    token: response.credential
  });

  if (error) {
    redirect('/login');
    return;
  }

  redirect('/');
}

export async function updateSession(request: NextRequest) {
  let supabaseResponse = NextResponse.next({
    request
  });

  const supabase = createServerClient(
    process.env.NEXT_PUBLIC_SUPABASE_URL!,
    process.env.NEXT_PUBLIC_SUPABASE_ANON_KEY!,
    {
      cookies: {
        getAll() {
          return request.cookies.getAll();
        },
        setAll(cookiesToSet) {
          cookiesToSet.forEach(({ name, value, options }) =>
            request.cookies.set(name, value)
          );
          supabaseResponse = NextResponse.next({
            request
          });
          cookiesToSet.forEach(({ name, value, options }) =>
            supabaseResponse.cookies.set(name, value, options)
          );
        }
      }
    }
  );

  const {
    data: { user }
  } = await supabase.auth.getUser();

  if (
    !user &&
    !request.nextUrl.pathname.startsWith('/login') &&
    !request.nextUrl.pathname.startsWith('/auth')
  ) {
    const url = request.nextUrl.clone();
    url.pathname = '/login';
    return NextResponse.redirect(url);
  }

  return supabaseResponse;
}
