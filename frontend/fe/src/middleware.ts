import { cookies } from "next/headers";
import { NextRequest, NextResponse } from "next/server"

export const middleware = async (req: NextRequest) => {
    const cookie = await cookies();
    const token = cookie.get("token");
    const isAuth = !!token;
    const authPage = req.nextUrl.pathname === "/";

    if (authPage) {
        if (isAuth) {
            return NextResponse.redirect(new URL("/home", req.url));
        }
        return null;
    }

    if (!isAuth) {
        return NextResponse.redirect(new URL("/", req.url));
    }
}

export const config = {
    matcher: ["/", "/home"]
}