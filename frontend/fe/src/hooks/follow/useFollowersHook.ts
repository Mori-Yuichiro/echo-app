import { RelationshipType } from "@/app/types/relationship";
import axiosInstance from "@/lib/axiosInstance"
import { useParams, useRouter } from "next/navigation";
import { useEffect, useState } from "react";

export const useFollowersHook = () => {
    const { instance } = axiosInstance();
    const router = useRouter();
    const [relationships, setRelationships] = useState<RelationshipType[] | null>(null);
    const { id } = useParams<{ id: string }>();

    useEffect(() => {
        const fetchData = async () => {
            try {
                const { data } = await instance.get<RelationshipType[]>(
                    `/users/${id}/followers`,
                    { withCredentials: true }
                );
                setRelationships(data);
            } catch (err) {
                console.error(err);
            }
        }
        fetchData();
    }, [])

    return {
        router,
        relationships
    };
}