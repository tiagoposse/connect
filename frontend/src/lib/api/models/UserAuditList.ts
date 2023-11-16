/* tslint:disable */
/* eslint-disable */
/**
 * Ent Schema API
 * This is an auto generated API description made out of an Ent schema definition
 *
 * The version of the OpenAPI document: 0.1.0
 * 
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
/**
 * 
 * @export
 * @interface UserAuditList
 */
export interface UserAuditList {
    /**
     * 
     * @type {string}
     * @memberof UserAuditList
     */
    id: string;
    /**
     * 
     * @type {string}
     * @memberof UserAuditList
     */
    action: string;
    /**
     * 
     * @type {string}
     * @memberof UserAuditList
     */
    author: string;
    /**
     * 
     * @type {Date}
     * @memberof UserAuditList
     */
    timestamp: Date;
}

/**
 * Check if a given object implements the UserAuditList interface.
 */
export function instanceOfUserAuditList(value: object): boolean {
    let isInstance = true;
    isInstance = isInstance && "id" in value;
    isInstance = isInstance && "action" in value;
    isInstance = isInstance && "author" in value;
    isInstance = isInstance && "timestamp" in value;

    return isInstance;
}

export function UserAuditListFromJSON(json: any): UserAuditList {
    return UserAuditListFromJSONTyped(json, false);
}

export function UserAuditListFromJSONTyped(json: any, ignoreDiscriminator: boolean): UserAuditList {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'id': json['id'],
        'action': json['action'],
        'author': json['author'],
        'timestamp': (new Date(json['timestamp'])),
    };
}

export function UserAuditListToJSON(value?: UserAuditList | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'id': value.id,
        'action': value.action,
        'author': value.author,
        'timestamp': (value.timestamp.toISOString()),
    };
}

