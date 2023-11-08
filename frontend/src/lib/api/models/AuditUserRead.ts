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
 * @interface AuditUserRead
 */
export interface AuditUserRead {
    /**
     * 
     * @type {string}
     * @memberof AuditUserRead
     */
    id: string;
    /**
     * 
     * @type {string}
     * @memberof AuditUserRead
     */
    email: string;
    /**
     * 
     * @type {string}
     * @memberof AuditUserRead
     */
    firstname: string;
    /**
     * 
     * @type {string}
     * @memberof AuditUserRead
     */
    lastname: string;
    /**
     * 
     * @type {string}
     * @memberof AuditUserRead
     */
    provider: string;
    /**
     * 
     * @type {string}
     * @memberof AuditUserRead
     */
    photoUrl?: string;
    /**
     * 
     * @type {boolean}
     * @memberof AuditUserRead
     */
    disabled: boolean;
    /**
     * 
     * @type {string}
     * @memberof AuditUserRead
     */
    disabledReason?: string;
    /**
     * 
     * @type {string}
     * @memberof AuditUserRead
     */
    groupId: string;
}

/**
 * Check if a given object implements the AuditUserRead interface.
 */
export function instanceOfAuditUserRead(value: object): boolean {
    let isInstance = true;
    isInstance = isInstance && "id" in value;
    isInstance = isInstance && "email" in value;
    isInstance = isInstance && "firstname" in value;
    isInstance = isInstance && "lastname" in value;
    isInstance = isInstance && "provider" in value;
    isInstance = isInstance && "disabled" in value;
    isInstance = isInstance && "groupId" in value;

    return isInstance;
}

export function AuditUserReadFromJSON(json: any): AuditUserRead {
    return AuditUserReadFromJSONTyped(json, false);
}

export function AuditUserReadFromJSONTyped(json: any, ignoreDiscriminator: boolean): AuditUserRead {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'id': json['id'],
        'email': json['email'],
        'firstname': json['firstname'],
        'lastname': json['lastname'],
        'provider': json['provider'],
        'photoUrl': !exists(json, 'photo_url') ? undefined : json['photo_url'],
        'disabled': json['disabled'],
        'disabledReason': !exists(json, 'disabled_reason') ? undefined : json['disabled_reason'],
        'groupId': json['group_id'],
    };
}

export function AuditUserReadToJSON(value?: AuditUserRead | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'id': value.id,
        'email': value.email,
        'firstname': value.firstname,
        'lastname': value.lastname,
        'provider': value.provider,
        'photo_url': value.photoUrl,
        'disabled': value.disabled,
        'disabled_reason': value.disabledReason,
        'group_id': value.groupId,
    };
}
